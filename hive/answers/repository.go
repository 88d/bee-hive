package answers

import (
	"github.com/black-banana/bee-hive/hive/users"
	"github.com/black-banana/bee-hive/rethink"
	r "github.com/dancannon/gorethink"
)

var (
	TableName       = "answers"
	QuestionIDField = "question_id"
	AuthorIDField   = "author_id"
)

// Repository for access to answers
type Repository struct {
	rethink.Repository
}

// NewRepository creates new repository for access to answers
func NewRepository() *Repository {
	var repo = &Repository{rethink.NewRepository(TableName)}
	rethink.CreateTableIndexIfNotExists(repo.Session, TableName, QuestionIDField)
	return repo
}

func (re *Repository) GetAll(questionID string) ([]*Answer, error) {
	res, err := re.Table().
		Filter(filterQuestionID(questionID)).
		Merge(mergeAuthor).
		Run(re.Session)
	defer res.Close()
	if err != nil {
		return nil, err
	}
	var answers = make([]*Answer, 0)
	return answers, res.All(&answers)
}

func (re *Repository) Create(questionID string, a *Answer) error {
	a.QuestionID = questionID
	res, err := re.Table().Insert(a).RunWrite(re.Session)
	if err != nil {
		return err
	}
	a.ID = res.GeneratedKeys[0]
	return nil
}

func (re *Repository) Update(questionID string, a *Answer) error {
	a.QuestionID = questionID
	_, err := re.Table().Filter(filterQuestionID(questionID)).Get(a.ID).Update(a).RunWrite(re.Session)
	return err
}

func (re *Repository) GetByID(questionID string, id string) (*Answer, error) {
	res, err := re.Table().GetAllByIndex(QuestionIDField, questionID).Get(id).Merge(mergeAuthor).Run(re.Session)
	defer res.Close()
	if err != nil {
		return nil, err
	}
	var a *Answer
	return a, res.One(&a)
}

func (re *Repository) RemoveByID(questionID string, id string) error {
	_, err := re.Table().Filter(filterQuestionID(questionID)).Get(id).Delete().RunWrite(re.Session)
	return err
}

func filterQuestionID(id string) map[string]interface{} {
	return map[string]interface{}{QuestionIDField: id}
}

func mergeAuthor(p r.Term) interface{} {
	return map[string]interface{}{
		AuthorIDField: r.Table(users.TableName).Get(p.Field(AuthorIDField)),
	}
}
