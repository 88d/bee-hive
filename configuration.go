package main

import (
	"log"

	"github.com/black-banana/bee-hive/hive/auth"
	"github.com/black-banana/bee-hive/rethink"
	"github.com/olebedev/config"
)

type Config struct {
	listen string
	db     *rethink.Config
	auth   *auth.Config
}

var defaultConfig = `{
		"listen":":9999",
		"db": {
			"server":"localhost:28015",
			"name":"beehive",
			"maxidle":10,
			"maxopen":10
		},
		"auth": {
			"expirehours":72
		}
	}`

var globalConfig = new(Config)

func loadConfiguration() {
	// Load Default Config
	cfg, err := config.ParseJson(defaultConfig)
	if err != nil {
		panic(err)
	}
	// Configuration
	cfgFile, err := config.ParseJsonFile("config.json")
	cfgFile.Env().Flag()
	if err != nil {
		log.Printf("no config file 'config.json' found will use default values")
	} else {
		cfg, err = cfg.Extend(cfgFile)
		if err != nil {
			panic(err)
		}
	}

	globalConfig.listen = getString(cfg, "listen")
	globalConfig.db = new(rethink.Config)
	globalConfig.db.Server = getString(cfg, "db.server")
	globalConfig.db.Name = getString(cfg, "db.name")
	globalConfig.db.MaxIdle = getInt(cfg, "db.maxidle")
	globalConfig.db.MaxOpen = getInt(cfg, "db.maxopen")
	globalConfig.auth = new(auth.Config)
	globalConfig.auth.ExpireHours = getInt(cfg, "auth.expirehours")
	globalConfig.auth.SigningKey = getString(cfg, "auth.signingkey")
}

func getString(cfg *config.Config, cfgValue string) string {
	var value, err = cfg.String(cfgValue)
	if err != nil {
		panic(cfgValue + " is Missing ")
	}
	log.Printf(cfgValue, value)
	return value
}

func getInt(cfg *config.Config, cfgValue string) int {
	var value, err = cfg.Int(cfgValue)
	if err != nil {
		panic(err)
	}
	log.Printf(cfgValue, value)
	return value
}
