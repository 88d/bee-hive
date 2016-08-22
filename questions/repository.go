package questions

import "github.com/black-banana/bee-hive/rethink"

type Repository struct {
	rethink.Repository
}

var repository *Repository

func NewRepository() *Repository {
	return &Repository{rethink.NewRepository("questions")}
}

func (re *Repository) GetAll() ([]Question, error) {
	res, err := re.Table().Run(re.Session)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	var questions []Question
	if err := res.All(&questions); err != nil {
		return nil, err
	}
	return questions, err
}

func (re *Repository) Create(q *Question) error {
	res, err := re.Table().Insert(q).RunWrite(re.Session)
	if err != nil {
		return err
	}
	q.ID = res.GeneratedKeys[0]
	return nil
}

func (re *Repository) Update(q *Question) error {
	if _, err := re.Table().Get(q.ID).Update(q).RunWrite(re.Session); err != nil {
		return err
	}
	return nil
}

func (re *Repository) RemoveByID(id string) error {
	if _, err := re.Table().Get(id).Delete().RunWrite(re.Session); err != nil {
		return err
	}
	return nil
}
