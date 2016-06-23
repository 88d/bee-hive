package rethink

import r "github.com/dancannon/gorethink"

type Repository struct {
	Session *r.Session
	table   string
}

func (re *Repository) Table() r.Term {
	return r.Table(re.table)
}

func (re *Repository) init() error {
	if err := CreateTableIfNotExists(re.Session, re.table); err != nil {
		return err
	}
	return nil
}

func NewRepository(table string) Repository {
	var repo = Repository{masterSession, table}
	if err := repo.init(); err != nil {
		panic(err)
	}
	return repo
}
