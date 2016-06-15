package main

import (
	"fmt"

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
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	var listen, listenMissing = cfg.String("listen")
	if listenMissing != nil {
		panic(listenMissing)
	}
	globalConfig.listen = listen
	var dbAddress, dbMissingServer = cfg.String("db.server")
	if dbMissingServer != nil {
		panic(dbMissingServer)
	}
	globalConfig.dbServer = dbAddress
	var dbName, dbMissingDatabase = cfg.String("db.name")
	if dbMissingDatabase != nil {
		panic(dbMissingDatabase)
	}
	globalConfig.dbName = dbName
}
