package util

import "fmt"

type Error struct {
	Location Location
	Message  string
}

type Errors []Error

func (e Errors) Format() string {
	whole := ""
	whole += fmt.Sprintf("%d errors\n", len(e))
	for _, err := range e {
		whole += fmt.Sprintf("%s at line %d: %s\n", err.Location.Filename, err.Location.Line, err.Message)
	}
	return whole
}
