package main

import (
	"log"

	"github.com/black-banana/bee-hive/rethink"
	"github.com/olebedev/config"
)

type Config struct {
	listen string
	db     *rethink.Config
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
	globalConfig.db = new(rethink.Config)
	globalConfig.db.Server = panicIfMissing(cfg, "db.server")
	globalConfig.db.Name = panicIfMissing(cfg, "db.name")
	globalConfig.db.MaxIdle = defaultIfMissingInt(cfg, "db.maxidle", 10)
	globalConfig.db.MaxOpen = defaultIfMissingInt(cfg, "db.maxopen", 10)
}

func panicIfMissing(cfg *config.Config, cfgValue string) string {
	var value, err = cfg.String(cfgValue)
	log.Printf(cfgValue, value)
	if err != nil {
		panic(err)
	}
	return value
}

func defaultIfMissingInt(cfg *config.Config, cfgValue string, defaultValue int) int {
	var value, err = cfg.Int(cfgValue)
	log.Printf(cfgValue, value)
	if err != nil {
		return defaultValue
	}
	return value
}
