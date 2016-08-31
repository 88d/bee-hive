package users

import (
	"github.com/black-banana/bee-hive/rethink"
	r "github.com/dancannon/gorethink"
	"github.com/juju/errors"
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
		return nil, errors.Annotate(err, "GetAll")
	}
	var users = make([]*User, 0)
	return users, errors.Annotate(res.All(&users), "GetAll")
}

func (re *Repository) Create(q *User) error {
	res, err := re.Table().Insert(q).RunWrite(re.Session)
	if err != nil {
		return errors.Annotate(err, "Create")
	}
	q.ID = res.GeneratedKeys[0]
	return nil
}

func (re *Repository) Update(q *User) error {
	_, err := re.Table().Get(q.ID).Update(q).RunWrite(re.Session)
	return errors.Annotate(err, "Update")
}

func (re *Repository) GetByID(id string) (*User, error) {
	res, err := re.Table().Get(id).Run(re.Session)
	defer res.Close()
	if err != nil {
		return nil, errors.Annotate(err, "GetByID")
	}
	var q *User
	return q, errors.Annotate(res.One(&q), "GetByID")
}

func (re *Repository) GetByName(name string) (*User, error) {
	res, err := re.Table().
		Filter(r.Row.Field("name").Eq(name)).Run(re.Session)
	defer res.Close()
	if err != nil {
		return nil, errors.Annotate(err, "GetByName")
	}
	var q *User
	return q, errors.Annotate(res.One(&q), "GetByName")
}

func (re *Repository) RemoveByID(id string) error {
	_, err := re.Table().Get(id).Delete().RunWrite(re.Session)
	return errors.Annotate(err, "RemoveByID")
}
