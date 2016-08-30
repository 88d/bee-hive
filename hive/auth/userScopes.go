package auth

type UserScopes struct {
	UserID string   `gorethink:"id,omitempty" json:"id,omitempty"`
	Scopes []string `gorethink:"scopes" json:"scopes"`
}
