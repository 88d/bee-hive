package rethink

import r "github.com/dancannon/gorethink"

// Repository is the base struct for all rethinkdb access
type Repository struct {
	Session *r.Session
	table   string
}

// TableName returns the table of this Repository
func (re *Repository) TableName() string {
	return re.table
}

// Table returns the gorethink table of this repository
func (re *Repository) Table() r.Term {
	return r.Table(re.table)
}

// CreateTableIndexIfNotExists is a shortcut for the same method
func (re *Repository) CreateTableIndexIfNotExists(fieldName string) error {
	return CreateTableIndexIfNotExists(re.Session, re.table, fieldName)
}

func (re *Repository) init() error {
	return CreateTableIfNotExists(re.Session, re.table)
}

// NewRepository creates a Repository with given table to access rethinkdb
func NewRepository(table string) Repository {
	var repo = Repository{masterSession, table}
	if err := repo.init(); err != nil {
		panic(err)
	}
	return repo
}
