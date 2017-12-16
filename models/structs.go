package models

type Commands []Command

type Command struct {
	Name        string
	Description string
	ArgsStr     string
	Args        []Arg
}

type Arg struct {
	Name        string
	Description string
}

