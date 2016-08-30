package answers

import (
	"time"

	"github.com/black-banana/bee-hive/hive/users"
)

// Answer stores information about answers to questions
type Answer struct {
	ID         string     `gorethink:"id,omitempty" json:"id,omitempty"`
	Content    string     `gorethink:"content" json:"content"`
	Author     users.User `gorethink:"author_id,reference" gorethink_ref:"id" json:"author,omitempty"`
	QuestionID string     `gorethink:"question_id" gorethink_ref:"id" json:"-"`
	CreatedAt  time.Time  `gorethink:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `gorethink:"update_at" json:"updated_at"`
}
