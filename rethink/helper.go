package rethink

import r "github.com/dancannon/gorethink"
import "log"

func isStringInArray(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func createDBIfNotExists(session *r.Session, dbName string) error {
	log.Println("CreateDBIfNotExists", dbName)
	dbListRes, err := r.DBList().Run(session)
	if err != nil {
		return err
	}
	var dbs []string
	if err := dbListRes.All(&dbs); err != nil {
		return err
	}
	if !isStringInArray(dbName, dbs) {
		log.Println("Create DB", dbName)
		errCreateDB := r.DBCreate(dbName).Exec(session)
		if errCreateDB != nil {
			return errCreateDB
		}
	}
	return nil
}

// CreateTableIfNotExists creates a table if this does not exist
func CreateTableIfNotExists(session *r.Session, table string) error {
	log.Println("CreateTableIfNotExists", table)
	tableListRes, err := r.TableList().Run(session)
	if err != nil {
		return err
	}
	var tables []string
	if err := tableListRes.All(&tables); err != nil {
		return err
	}
	if !isStringInArray(table, tables) {
		log.Println("Create table", table)
		errCreateDB := r.TableCreate(table).Exec(session)
		if errCreateDB != nil {
			return err
		}
	}
	return nil
}

func CreateTableIndexIfNotExists(session *r.Session, table string, index string) error {
	log.Println("CreateTableIndexIfNotExists", table)
	indexListRes, err := r.Table(table).IndexList().Run(session)
	if err != nil {
		return err
	}
	var indexes []string
	if err := indexListRes.All(&indexes); err != nil {
		return err
	}
	if !isStringInArray(index, indexes) {
		log.Println("Create index", index)
		errCreateDB := r.Table(table).IndexCreate(index).Exec(session)
		if errCreateDB != nil {
			return err
		}
	}
	return nil
}

var masterSession *r.Session

// StartMasterSession must be called to connect to the server
func StartMasterSession(config *Config) {
	var err error
	masterSession, err = r.Connect(r.ConnectOpts{
		Address:  config.Server,
		Database: config.Name,
		MaxIdle:  config.MaxIdle,
		MaxOpen:  config.MaxOpen,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err := createDBIfNotExists(masterSession, config.Name); err != nil {
		log.Fatalln(err.Error())
	}
}

// StopMasterSession should be called after the application stops
func StopMasterSession() {
	masterSession.Close()
}
