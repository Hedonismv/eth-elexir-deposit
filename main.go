package main

import (
	"context"
	"elexir/actions"
	"elexir/config"
	"elexir/delayer"
	"elexir/formatter"
	"elexir/stateManager"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	"io"
	"os"
	"time"
)

func loadConfig() (*config.Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Error reading config file, %s", err)
	}

	var config config.Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("Unable to decode into struct, %v", err)
	}

	return &config, nil
}

var startAppLogger = log.NewWithOptions(os.Stdout, log.Options{
	ReportTimestamp: true,
	TimeFormat:      time.Kitchen,
	Prefix:          "Start App 🚀",
})

func main() {
	logFile, err := os.OpenFile("elexir.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	var logger = log.NewWithOptions(multiWriter, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "Main 💻",
	})

	loadedConfig, err := loadConfig()
	if err != nil {
		logger.Fatalf("Failed to load loadedConfig: %v", err)
	}

	startAppLogger.Infof("App Name: %s", loadedConfig.App.Name)
	startAppLogger.Infof("App Version: %s", loadedConfig.App.Version)
	startAppLogger.Infof("Rpc Provider: %s", loadedConfig.Ethereum.Rpc)
	startAppLogger.Infof("Delays between wallets (Seconds) Min:%d / Max:%d", loadedConfig.Ethereum.Delays.Wallet.Min, loadedConfig.Ethereum.Delays.Wallet.Max)
	startAppLogger.Infof("Delays between blocks (Seconds) Min:%d / Max:%d", loadedConfig.Ethereum.Delays.Block.Min, loadedConfig.Ethereum.Delays.Block.Max)
	startAppLogger.Infof("Work Amount Range (Percent) Min:%d / Max:%d", loadedConfig.Ethereum.Workflow.WorkAmountRangePercent.Min, loadedConfig.Ethereum.Workflow.WorkAmountRangePercent.Max)
	startAppLogger.Infof("Gas Limit (Gwei): %d", loadedConfig.Ethereum.Workflow.GweiLimit)
	startAppLogger.Info("Creator: Vi Hedo")
	startAppLogger.Info("X(Twitter): https://x.com/hedonismatall\n")

	client, err := ethclient.Dial(loadedConfig.Ethereum.Rpc)
	if err != nil {
		logger.Errorf("Failed to connect to the Ethereum client: %v", err)
	}

	logger.Info("Connected to Ethereum client")

	keys, err := formatter.ReadKeysFromFile("keys.txt")
	if err != nil {
		logger.Fatalf("Error reading keys from file: %v", err)
	}

	logger.Infof("Loaded %v keys", len(keys))

	stateFilePath := "state.json"
	state, err := stateManager.LoadState(stateFilePath)
	if err != nil {
		logger.Fatalf("Failed to load state: %v", err)
		return
	}

	//! Main Loop
	for _, key := range keys {
		hex256Pk := formatter.PrivateKeyToHex(key)

		privateKeyECDSA, err := crypto.HexToECDSA(hex256Pk)
		if err != nil {
			logger.Fatalf("Failed to parse private key: %v", err)
		}
		fromAddress := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)

		//! checking status
		if state[fromAddress.Hex()] {
			logger.Infof("Skipping address: %s (already done)", fromAddress.Hex())
			continue
		}

		logger.Infof("Working with address: %s", fromAddress.Hex())

		//! Eth Balance
		balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
		if err != nil {
			logger.Errorf("Failed to get balance: %v", err)
			continue
		}

		//! Generating random amount of Eth for deposit to Elixir
		amount := formatter.GetRandomAmount(balance, loadedConfig.Ethereum.Workflow.WorkAmountRangePercent.Min, loadedConfig.Ethereum.Workflow.WorkAmountRangePercent.Max)
		logger.Infof("Generated amount: %v ETH", formatter.ConvertWeiToEther(amount))

		//! Deposit to Elixir
		logger.Infof("Depositing %v ETH to Elixir", formatter.ConvertWeiToEther(amount))
		actions.DepositToElixir(privateKeyECDSA, client, amount, loadedConfig)

		state[fromAddress.Hex()] = true
		err = stateManager.SaveState(stateFilePath, state)
		if err != nil {
			logger.Errorf("Failed to save state: %v", err)
		}

		//! Delay Wallets
		delayer.DelayWallet(loadedConfig)
	}

}
