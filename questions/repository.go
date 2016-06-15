package questions

import r "github.com/dancannon/gorethink"

type Repository struct {
	session  *r.Session
	database string
	table    string
}

func (re Repository) Table() r.Term {
	return r.DB(re.database).Table(re.table)
}

func NewRepository(server string, database string) Repository {
	var session *r.Session
	session, err := r.Connect(r.ConnectOpts{
		Address: server,
	})
	if err != nil {
		panic(err)
	}
	return Repository{session, database, "questions"}
}

func (re Repository) GetAll() ([]Question, error) {
	res, err := re.Table().Run(re.session)
	if err != nil {
		return nil, err
	}
	var questions []Question
	err = res.All(&questions)
	if err != nil {
		return nil, err
	}
	return questions, err
}

func (re Repository) Create(q *Question) error {
	res, err := re.Table().Insert(q).RunWrite(re.session)
	if err != nil {
		return err
	}
	q.ID = res.GeneratedKeys[0]
	return nil
}

func (re Repository) Update(q *Question) error {
	if _, err := re.Table().Get(q.ID).Update(q).RunWrite(re.session); err != nil {
		return err
	}
	return nil
}

func (re Repository) Remove(id string) error {
	if _, err := re.Table().Get(id).Delete().RunWrite(re.session); err != nil {
		return err
	}
	return nil
}
