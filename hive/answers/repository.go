package answers

import (
	"github.com/black-banana/bee-hive/hive/users"
	"github.com/black-banana/bee-hive/rethink"
	r "github.com/dancannon/gorethink"
	"github.com/juju/errors"
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
		return nil, errors.Annotate(err, "GetAll")
	}
	var answers = make([]*Answer, 0)
	return answers, errors.Annotate(res.All(&answers), "GetAll")
}

func (re *Repository) Create(questionID string, a *Answer) error {
	a.QuestionID = questionID
	res, err := re.Table().Insert(a).RunWrite(re.Session)
	if err != nil {
		return errors.Annotate(err, "Create")
	}
	a.ID = res.GeneratedKeys[0]
	return nil
}

func (re *Repository) Update(questionID string, a *Answer) error {
	a.QuestionID = questionID
	_, err := re.Table().Filter(filterQuestionID(questionID)).Get(a.ID).Update(a).RunWrite(re.Session)
	return errors.Annotate(err, "Update")
}

func (re *Repository) GetByID(questionID string, id string) (*Answer, error) {
	res, err := re.Table().GetAllByIndex(QuestionIDField, questionID).Get(id).Merge(mergeAuthor).Run(re.Session)
	defer res.Close()
	if err != nil {
		return nil, err
	}
	var a *Answer
	return a, errors.Annotate(res.One(&a), "GetByID")
}

func (re *Repository) RemoveByID(questionID string, id string) error {
	_, err := re.Table().Filter(filterQuestionID(questionID)).Get(id).Delete().RunWrite(re.Session)
	return errors.Annotate(err, "RemoveByID")
}

func filterQuestionID(id string) map[string]interface{} {
	return map[string]interface{}{QuestionIDField: id}
}

func mergeAuthor(p r.Term) interface{} {
	return map[string]interface{}{
		AuthorIDField: r.Table(users.TableName).Get(p.Field(AuthorIDField)),
	}
}
