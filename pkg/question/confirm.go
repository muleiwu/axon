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
	if strings.Contains(question, "</>") {
		color.Print(question)
	} else {
		color.Info.Print(question)
	}
	color.Info.Print(" [")
	color.LightGreen.Print("Y")
	color.Info.Print("es/")
	color.LightRed.Print("N")
	color.Info.Print("o] : ")
	input, _ := q.readInput()
	return q.validateInput(input, def.(bool))
}

func (q *Confirm) readInput() (string, error) {
	var input string
	_, err := fmt.Scanln(&input)
	color.Println()
	return input, err
}

func (q *Confirm) validateInput(input string, defaultValue bool) any {
	if input == "" {
		return defaultValue
	}
	switch strings.ToLower(input) {
	case "y":
		return true
	case "yes":
		return true
	case "n":
		return false
	case "no":
		return false
	default:
		return q.Execute("<warn>无效输入，请重新输入</>", defaultValue)
	}
}
