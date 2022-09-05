package monster

import "os"

type Monster struct {
	Name string
}

func New() *Monster {
	monster := os.Getenv("monster")
	return &Monster{Name: monster}
}
