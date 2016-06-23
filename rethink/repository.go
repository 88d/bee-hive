package rethink

import r "github.com/dancannon/gorethink"

type Repository struct {
	Session *r.Session
	table   string
}

func (re Repository) Table() r.Term {
	return r.Table(re.table)
}

func initRepository(repo Repository) error {
	if err := CreateTableIfNotExists(repo.Session, repo.table); err != nil {
		return err
	}
	return nil
}

func NewRepository(table string) Repository {
	var repo = Repository{masterSession, table}
	if err := initRepository(repo); err != nil {
		panic(err)
	}
	return repo
}
