package actions

import (
	"context"
	"crypto/ecdsa"
	"elexir/config"
	"elexir/elixirContract"
	"elexir/formatter"
	"elexir/gasGuard"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"math/big"
	"os"
	"time"
)

func DepositToElixir(privateKeyECDSA *ecdsa.PrivateKey, client *ethclient.Client, ethAmount *big.Int, cfg *config.Config) {

	logFile, err := os.OpenFile("elexir.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	var logger = log.NewWithOptions(multiWriter, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "Elixir 🍺",
	})

	fromAddress := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)

	elixirContractAddress := common.HexToAddress("0x1F75881DC0707b5236f739b5B64A87c211294Abb")

	ctx := context.Background()

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		logger.Errorf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		logger.Errorf("Failed to get gas price: %v", err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		logger.Errorf("Failed to get chain ID: %v", err)
	}

	//! ABIGEN INSTANCE
	elixirInstance, err := elixirContract.NewElixirContract(elixirContractAddress, client)
	if err != nil {
		logger.Errorf("Failed to instantiate a Token contract: %v", err)
	}

	contractABI, err := elixirContract.ElixirContractMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to pack function input: %v", err)
	}

	callData, err := contractABI.Pack("deposit")
	if err != nil {
		logger.Errorf("Failed to pack calldata: %v", err)
	}

	gasGuard.CheckGasPrice(client, cfg)

	gasLimit, err := client.EstimateGas(ctx, ethereum.CallMsg{
		To:    &elixirContractAddress,
		Data:  callData,
		Value: ethAmount,
	})

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainID)
	if err != nil {
		logger.Errorf("Failed to create auth: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = ethAmount
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	logger.Infof("Depositing %v ETH to Elixir", formatter.ConvertWeiToEther(ethAmount))

	tx, err := elixirInstance.Deposit(auth)
	if err != nil {
		logger.Fatalf("Failed to deposit: %v", err)
	}
	scan := "https://etherscan.io/tx/" + tx.Hash().Hex()
	logger.Infof("Deposit pending: %s", scan)

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		logger.Fatalf("Failed to wait for tx: %v", err)
	}
	link := "https://etherscan.io/tx/" + receipt.TxHash.Hex()

	logger.Infof("Deposit to Elixir successful: %s", link)
}
