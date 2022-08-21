package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"time"
)

type ConfigList struct {
	ApiKey      string
	ApiSecret   string
	LogFile     string
	ProductCode string

	TradeDuration time.Duration
	Durations     map[string]time.Duration
	DbName        string
	SQLDriver     string
	Port          int
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Faild to read file: %v", err)
		os.Exit(1)
	}

	durations := map[string]time.Duration{
		"ls": time.Second,
		"lm": time.Minute,
		"lh": time.Hour,
	}

	Config = ConfigList{
		ApiKey:        cfg.Section("BitFlyer").Key("api_key").String(),
		ApiSecret:     cfg.Section("BitFlyer").Key("api_secret").String(),
		LogFile:       cfg.Section("gotrading").Key("log_file").String(),
		ProductCode:   cfg.Section("gotrading").Key("product_code").String(),
		Durations:     durations,
		TradeDuration: durations[cfg.Section("gotrading").Key("trade_duration").String()],
		DbName:        cfg.Section("db").Key("name").String(),
		SQLDriver:     cfg.Section("db").Key("driver").String(),
		Port:          cfg.Section("web").Key("port").MustInt(),
	}

}
