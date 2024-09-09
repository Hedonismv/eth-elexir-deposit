package gasGuard

import (
	"context"
	"elexir/config"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
	"time"
)

func CheckGasPrice(client *ethclient.Client, cfg *config.Config) {

	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "Gas Guard ⛽",
	})
	limit := big.NewInt(int64(cfg.Ethereum.Workflow.GweiLimit))
	for {
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			logger.Errorf("Failed to get gas price: %v", err)
		}

		gasPriceGwei := new(big.Int).Div(gasPrice, big.NewInt(1e9))
		logger.Infof("Current gas price: %s Gwei | Limit is: %s", gasPriceGwei.String(), limit.String())

		if gasPriceGwei.Cmp(limit) <= 0 {
			logger.Infof("Gas price is within the limit, proceeding...")
			break
		}

		logger.Warnf("Gas price is too high, waiting...")
		time.Sleep(30 * time.Second)
	}
}
