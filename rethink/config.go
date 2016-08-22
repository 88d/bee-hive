package rethink

type Config struct {
	Server  string
	Name    string
	MaxIdle int
	MaxOpen int
}
