package main

import (
	"gotrading/bitflyer"
	"gotrading/config"
	"gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	tickerChannel := make(chan bitflyer.Ticker)
	apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
}
