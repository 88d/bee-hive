package question

import "github.com/black-banana/bee-hive/hive/user"
import "github.com/black-banana/bee-hive/hive/answer"

type Question struct {
	ID          string          `gorethink:"id,omitempty" json:"id,omitempty"`
	Title       string          `gorethink:"title" json:"title,omitempty"`
	Description string          `gorethink:"description" json:"description,omitempty"`
	Author      user.User       `gorethink:"author_id,reference,omitempty" gorethink_ref:"id" json:"author,omitempty"`
	Answers     []answer.Answer `gorethink:"answers,omitempty" json:"answers,omitempty"`
	Tags        []string        `gorethink:"tags,omitempty" json:"tags,omitempty"`
}
