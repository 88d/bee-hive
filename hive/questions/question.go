package questions

import (
	"time"

	"github.com/black-banana/bee-hive/hive/answers"
	"github.com/black-banana/bee-hive/hive/users"
)

// Question is used to store information about a question
type Question struct {
	ID          string           `gorethink:"id,omitempty" json:"id,omitempty"`
	Title       string           `gorethink:"title,omitempty" json:"title,omitempty"`
	Description string           `gorethink:"description,omitempty" json:"description,omitempty"`
	Author      users.User       `gorethink:"author_id,reference,omitempty" gorethink_ref:"id" json:"author,omitempty"`
	Answers     []answers.Answer `gorethink:"answers,omitempty" json:"answers,omitempty"`
	Tags        []string         `gorethink:"tags,omitempty" json:"tags,omitempty"`
	CreatedAt   time.Time        `gorethink:"created_at" json:"created_at"`
	UpdatedAt   time.Time        `gorethink:"update_at" json:"updated_at"`
}
