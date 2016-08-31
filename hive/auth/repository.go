package auth

import (
	"github.com/black-banana/bee-hive/rethink"
	"github.com/juju/errors"
)

var (
	TableName = "user_scopes"
	UserID    = "id"
)

// Repository for access to answers
type Repository struct {
	rethink.Repository
}

// NewRepository creates new repository for access to answers
func NewRepository() *Repository {
	var repo = &Repository{rethink.NewRepository(TableName)}
	repo.CreateTableIndexIfNotExists(UserID)
	return repo
}

func (re *Repository) GetUserScopes(userID string) (*UserScopes, error) {
	res, err := re.Table().
		Run(re.Session)
	defer res.Close()
	if err != nil {
		return nil, err
	}
	var userScopes UserScopes
	return &userScopes, errors.Annotate(res.One(&userScopes), "GetUserScopes")
}

func (re *Repository) CreateUserScopes(userID string) (*UserScopes, error) {
	var userScopes = &UserScopes{userID, []string{}}
	_, err := re.Table().Insert(userScopes).RunWrite(re.Session)
	return userScopes, errors.Annotate(err, "CreateUserScopes")
}

func (re *Repository) UpdateUserScopes(a *UserScopes) error {
	_, err := re.Table().Get(a.UserID).Update(a).RunWrite(re.Session)
	return errors.Annotate(err, "UpdateUserScopes")
}

func (re *Repository) RevokeScope(userID string, scope string) error {
	userScopes, err := re.GetUserScopes(userID)
	if err != nil {
		return errors.Annotate(err, "RevokeScope")
	}
	userScopes.Scopes = deleteStringInArray(userScopes.Scopes, scope)
	return errors.Annotate(re.UpdateUserScopes(userScopes), "RevokeScope")
}

func (re *Repository) GrantScope(userID string, scope string) error {
	userScopes, err := re.GetUserScopes(userID)
	if err != nil {
		return errors.Annotate(err, "GrantScope")
	}
	userScopes.Scopes = append(userScopes.Scopes, scope)
	return errors.Annotate(re.UpdateUserScopes(userScopes), "GrantScope")
}

func isStringInArray(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func indexOf(list []string, search string) int {
	for i, b := range list {
		if b == search {
			return i
		}
	}
	return -1
}

func deleteStringInArray(list []string, element string) []string {
	var i = indexOf(list, element)
	if i == -1 {
		return list
	}
	copy(list[i:], list[i+1:])
	list[len(list)-1] = "" // or the zero value of T
	list = list[:len(list)-1]
	return list
}
