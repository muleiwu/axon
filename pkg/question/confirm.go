package question

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

type Confirm struct {
}

func NewConfirm() *Confirm {
	return &Confirm{}
}

func (q *Confirm) Execute(question string, def any) any {
	color.Print(question)
	input, _ := q.readInput()
	return q.validateInput(input, def.(bool))
}

func (q *Confirm) readInput() (string, error) {
	var input string
	_, err := fmt.Scanln(&input)
	return input, err
}

func (q *Confirm) validateInput(input string, def bool) any {
	if input == "" {
		return def
	}
	switch strings.ToLower(input) {
	case "y":
		return true
	case "n":
		return false
	default:
		return q.Execute("\n<warn>无效输入，请重新输入 (y/n): </>", def)
	}
}
