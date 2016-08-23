package users

import "github.com/black-banana/bee-hive/rethink"

var TableName = "users"

type Repository struct {
	rethink.Repository
}

var repository *Repository

func NewRepository() *Repository {
	return &Repository{rethink.NewRepository(TableName)}
}

func (re *Repository) GetAll() ([]*User, error) {
	res, err := re.Table().Run(re.Session)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	var users = make([]*User, 0)
	if err := res.All(&users); err != nil {
		return nil, err
	}
	return users, err
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
	if _, err := re.Table().Get(q.ID).Update(q).RunWrite(re.Session); err != nil {
		return err
	}
	return nil
}

func (re *Repository) GetByID(id string) (*User, error) {
	res, err := re.Table().Get(id).Run(re.Session)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	var q *User
	if err := res.One(&q); err != nil {
		return nil, err
	}
	return q, nil
}

func (re *Repository) RemoveByID(id string) error {
	if _, err := re.Table().Get(id).Delete().RunWrite(re.Session); err != nil {
		return err
	}
	return nil
}
