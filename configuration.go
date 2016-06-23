package main

import (
	"log"

	"github.com/olebedev/config"
)

type Config struct {
	listen   string
	dbServer string
	dbName   string
}

var globalConfig = new(Config)

func LoadConfiguration() {
	// Configuration
	cfg, err := config.ParseJsonFile("config.json")
	cfg.Env().Flag()
	if err != nil {
		panic(err)
	}
	globalConfig.listen = panicIfMissing(cfg, "listen")
	globalConfig.dbServer = panicIfMissing(cfg, "db.server")
	globalConfig.dbName = panicIfMissing(cfg, "db.name")
}

func panicIfMissing(cfg *config.Config, cfgValue string) string {
	var value, err = cfg.String(cfgValue)
	log.Printf(cfgValue, value)
	if err != nil {
		panic(err)
	}
	return value
}
