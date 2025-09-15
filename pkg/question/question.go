package question

import (
	"fmt"

	"github.com/gookit/color"
)

type Question struct{}

func NewQuestion() *Question {
	return &Question{}
}

func (q *Question) Execute(question string, def any) any {
	color.Print(question)
	input, _ := q.readInput()
	return q.validateInput(input, def.(string))
}

func (q *Question) readInput() (string, error) {
	var input string
	_, err := fmt.Scanln(&input)
	return input, err
}

func (q *Question) validateInput(input string, def string) string {
	if input == "" {
		return def
	}

	return input
}
