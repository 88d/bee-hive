package questions

type Question struct {
	ID          string `gorethink:"id,omitempty" json:"id,omitempty"`
	Title       string `gorethink:"title" json:"title,omitempty"`
	Description string `gorethink:"description" json:"description,omitempty"`
	Author      User   `gorethink:"author_id,reference" gorethink_ref:"id" json:"author,omitempty"`
}

type User struct {
	ID   string `gorethink:"id,omitempty" json:"id,omitempty"`
	Name string `gorethink:"name,omitempty" json:"name,omitempty"`
}
