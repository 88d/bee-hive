package answer

import "github.com/black-banana/bee-hive/rethink"
import r "github.com/dancannon/gorethink"

var TableName = "answers"
var QuestionIndex = "question_id"

type Repository struct {
	rethink.Repository
}

var repository *Repository

func NewRepository() *Repository {
	var repo = &Repository{rethink.NewRepository(TableName)}
	rethink.CreateTableIndexIfNotExists(repo.Session, TableName, "question_id")
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
	if err := res.All(&answers); err != nil {
		return nil, err
	}
	return answers, err
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
	if _, err := re.Table().Filter(filterQuestionID(questionID)).Get(a.ID).Update(a).RunWrite(re.Session); err != nil {
		return err
	}
	return nil
}

func (re *Repository) GetByID(questionID string, id string) (*Answer, error) {
	res, err := re.Table().GetAllByIndex("question_id", questionID).Get(id).Merge(mergeAuthor).Run(re.Session)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	var a *Answer
	if err := res.One(&a); err != nil {
		return nil, err
	}
	return a, nil
}

func (re *Repository) RemoveByID(questionID string, id string) error {
	if _, err := re.Table().Filter(filterQuestionID(questionID)).Get(id).Delete().RunWrite(re.Session); err != nil {
		return err
	}
	return nil
}

func filterQuestionID(id string) map[string]interface{} {
	return map[string]interface{}{"question_id": id}
}

func mergeAuthor(p r.Term) interface{} {
	return map[string]interface{}{
		"author_id": r.Table("users").Get(p.Field("author_id")),
	}
}
