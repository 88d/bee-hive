package users

import (
	"github.com/black-banana/bee-hive/rethink"
	r "github.com/dancannon/gorethink"
)

var TableName = "users"

// Repository for access to users
type Repository struct {
	rethink.Repository
}

// NewRepository creates new repository for access to users
func NewRepository() *Repository {
	return &Repository{rethink.NewRepository(TableName)}
}

func (re *Repository) GetAll() ([]*User, error) {
	res, err := re.Table().Run(re.Session)
	defer res.Close()
	if err != nil {
		return nil, err
	}
	var users = make([]*User, 0)
	return users, res.All(&users)
}

func (re *Repository) Create(q *User) error {
	res, err := re.Table().Insert(q).RunWrite(re.Session)
	if err != nil {
		return err
	}
	q.ID = res.GeneratedKeys[0]
	return nil
}

func (re *Repository) Update(q *User) error {
	_, err := re.Table().Get(q.ID).Update(q).RunWrite(re.Session)
	return err
}

func (re *Repository) GetByID(id string) (*User, error) {
	res, err := re.Table().Get(id).Run(re.Session)
	defer res.Close()
	if err != nil {
		return nil, err
	}
	var q *User
	return q, res.One(&q)
}

func (re *Repository) GetByName(name string) (*User, error) {
	res, err := re.Table().
		Filter(r.Row.Field("name").Eq(name)).Run(re.Session)
	defer res.Close()
	if err != nil {
		return nil, err
	}
	var q *User
	return q, res.One(&q)
}

func (re *Repository) RemoveByID(id string) error {
	_, err := re.Table().Get(id).Delete().RunWrite(re.Session)
	return err
}
