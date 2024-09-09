package delayer

import (
	"elexir/config"
	"github.com/charmbracelet/log"
	"math/rand"
	"os"
	"time"
)

var logger = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: true,
	TimeFormat:      time.Kitchen,
	Prefix:          "Sleeping 😴",
})

func DelayBlock(config *config.Config) {
	blockDelay := rand.Intn(config.Ethereum.Delays.Block.Max-config.Ethereum.Delays.Block.Min) + config.Ethereum.Delays.Block.Min
	logger.Infof("Waiting for %d seconds\n", blockDelay)
	time.Sleep(time.Duration(blockDelay) * time.Second)
}

func DelayWallet(config *config.Config) {
	walletDelay := rand.Intn(config.Ethereum.Delays.Wallet.Max-config.Ethereum.Delays.Wallet.Min) + config.Ethereum.Delays.Wallet.Min
	logger.Infof("Waiting for %d seconds\n", walletDelay)
	time.Sleep(time.Duration(walletDelay) * time.Second)
}
