package questions

import r "github.com/dancannon/gorethink"

import "log"

type Repository struct {
	session  *r.Session
	database string
	table    string
}

func (re Repository) Table() r.Term {
	return r.DB(re.database).Table(re.table)
}

func initDB(session *r.Session, dbName string) error {
	log.Println("initDB", dbName)
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
		_, errCreateDB := r.DBCreate(dbName).RunWrite(session)
		if errCreateDB != nil {
			return errCreateDB
		}
	}
	return nil
}

func initTable(session *r.Session, dbName string, table string) error {
	log.Println("initTable", table, "in DB", dbName)
	tableListRes, err := r.DB(dbName).TableList().Run(session)
	if err != nil {
		return err
	}
	var tables []string
	if err := tableListRes.All(&tables); err != nil {
		return err
	}
	if !isStringInArray(table, tables) {
		log.Println("Create table", table, "in DB", dbName)
		_, errCreateDB := r.DB(dbName).TableCreate(table).RunWrite(session)
		if errCreateDB != nil {
			panic(errCreateDB)
		}
	}
	return nil
}

func initRepository(repo Repository) error {
	if err := initDB(repo.session, repo.database); err != nil {
		return err
	}
	if err := initTable(repo.session, repo.database, repo.table); err != nil {
		return err
	}
	return nil
}

func isStringInArray(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func NewRepository(server string, database string) Repository {
	var session *r.Session
	session, err := r.Connect(r.ConnectOpts{
		Address: server,
	})
	if err != nil {
		panic(err)
	}
	var repo = Repository{session, database, "questions"}
	if err := initRepository(repo); err != nil {
		panic(err)
	}
	return repo
}

func (re Repository) GetAll() ([]Question, error) {
	re.session.Query(r.Query{})
	res, err := re.Table().Run(re.session)
	if err != nil {
		return nil, err
	}
	var questions []Question
	err = res.All(&questions)
	if err != nil {
		return nil, err
	}
	return questions, err
}

func (re Repository) Create(q *Question) error {
	res, err := re.Table().Insert(q).RunWrite(re.session)
	if err != nil {
		return err
	}
	q.ID = res.GeneratedKeys[0]
	return nil
}

func (re Repository) Update(q *Question) error {
	if _, err := re.Table().Get(q.ID).Update(q).RunWrite(re.session); err != nil {
		return err
	}
	return nil
}

func (re Repository) Remove(id string) error {
	if _, err := re.Table().Get(id).Delete().RunWrite(re.session); err != nil {
		return err
	}
	return nil
}
