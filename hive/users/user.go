package users

type User struct {
	ID   string `gorethink:"id,omitempty" json:"id,omitempty"`
	Name string `gorethink:"name,omitempty" json:"name,omitempty"`
}
