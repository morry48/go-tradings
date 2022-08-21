package main

import (
	"fmt"
	"gotrading/app/controllers"
	"gotrading/app/models"
	"gotrading/config"
	"gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
	controllers.StreamIngestionData()
	//apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

	//tickerChannel := make(chan bitflyer.Ticker)
	//apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
}
