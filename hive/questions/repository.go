package questions

import "github.com/black-banana/bee-hive/rethink"
import r "github.com/dancannon/gorethink"

var TableName = "questions"

type Repository struct {
	rethink.Repository
}

var repository *Repository

func NewRepository() *Repository {
	return &Repository{rethink.NewRepository(TableName)}
}

func (re *Repository) GetAll() ([]*Question, error) {
	res, err := re.Table().
		Merge(mergeAuthor).Run(re.Session)
	defer res.Close()
	if err != nil {
		return nil, err
	}
	var questions = make([]*Question, 0)
	return questions, res.All(&questions)
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
	_, err := re.Table().Get(q.ID).Update(q).RunWrite(re.Session)
	return err
}

func (re *Repository) GetByID(id string) (*Question, error) {
	res, err := re.Table().
		Get(id).
		Default(defaults).
		Merge(mergeAuthor).
		Merge(mergeAnswers(id)).
		Run(re.Session)
	defer res.Close()
	if err != nil {
		return nil, err
	}
	var q *Question
	return q, res.One(&q)
}

func (re *Repository) RemoveByID(id string) error {
	if _, err := re.Table().Get(id).Delete().RunWrite(re.Session); err != nil {
		return err
	}
	return nil
}

var defaults = map[string]interface{}{
	"author_id": "",
}

func mergeAuthor(p r.Term) interface{} {
	return map[string]interface{}{
		"author_id": r.Table("users").Get(p.Field("author_id")),
	}
}

func mergeAnswers(id string) func(r.Term) interface{} {
	return func(p r.Term) interface{} {
		return map[string]interface{}{
			"answers": r.Table("answers").
				GetAllByIndex("question_id", id).
				Merge(mergeAuthor).
				CoerceTo("array"),
		}
	}
}
