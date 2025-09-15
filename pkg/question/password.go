package question

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/gookit/color"
	"golang.org/x/term"
)

type Password struct {
}

func NewPassword() *Password {
	return &Password{}
}

func (q *Password) Execute(question string, def any) any {
	color.Print(question)
	input, err := q.readInput()
	if err != nil {
		color.Error.Printf("\n%s^C\n", err.Error())
		os.Exit(0)
	}

	return q.validateInput(input, def.(string))
}

func (q *Password) validateInput(input string, def string) string {
	if input == "" {
		input = def
	}
	color.Println("")
	return input
}

func (q *Password) readInput() (string, error) {
	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	defer term.Restore(int(os.Stdin.Fd()), state)

	reader := bufio.NewReader(os.Stdin)
	var password []byte

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		switch r {
		case '\r', '\n': // 回车或换行
			fmt.Println() // 换行
			return string(password), nil
		case '\b', 127: // 退格
			if len(password) > 0 {
				password = password[:len(password)-1]
				fmt.Print("\b \b") // 删除最后一个字符并替换为空格
			}
		case 3: // Ctrl + C
			return "", errors.New("操作已取消")
		default:
			password = append(password, byte(r))
			fmt.Print("*") // 显示星号
		}
	}

	return string(password), nil
}
