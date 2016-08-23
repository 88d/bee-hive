package rethink

// Config used to configure rethinkdb master session
type Config struct {
	Server  string
	Name    string
	MaxIdle int
	MaxOpen int
}
