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
	var listen, listenMissing = cfg.String("listen")
	if listenMissing != nil {
		panic(listenMissing)
	}
	globalConfig.listen = listen
	log.Println("Listen", listen)
	var dbServer, dbMissingServer = cfg.String("db.server")
	if dbMissingServer != nil {
		panic(dbMissingServer)
	}
	globalConfig.dbServer = dbServer
	log.Println("DBServer", dbServer)
	var dbName, dbMissingDatabase = cfg.String("db.name")
	if dbMissingDatabase != nil {
		panic(dbMissingDatabase)
	}
	globalConfig.dbName = dbName
	log.Println("DBName", dbName)
}
