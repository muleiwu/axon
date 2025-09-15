package question

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/gookit/color"
	"golang.org/x/term"
)

type PhoneNumber struct{}

func NewPhoneNumber() *PhoneNumber {
	return &PhoneNumber{}
}

// Execute 执行电话号码询问流程
func (q *PhoneNumber) Execute(question string, def any) any {
	color.Print(question)
	input, err := q.readInput()
	if err != nil {
		color.Warnln(err.Error())
		os.Exit(0)
	}

	return q.validateInput(input, def.(string))
}

func (q *PhoneNumber) readInput() (string, error) {
	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	defer term.Restore(int(os.Stdin.Fd()), state)

	reader := bufio.NewReader(os.Stdin)
	var inputData []byte

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		switch r {
		case '\r', '\n': // 回车或换行
			maskedNumber := q.maskPhoneNumber(inputData)
			q.clearPreviousInput(len(inputData))
			q.println(fmt.Sprintf("%s（为保护您的隐私已脱敏）", maskedNumber))
			return string(inputData), nil
		case '\b', 127: // 退格
			if len(inputData) > 0 {
				inputData = inputData[:len(inputData)-1]
				fmt.Print("\b \b") // 删除最后一个字符并替换为空格
			}
		case 3: // Ctrl + C
			return "", errors.New("操作已取消")
		default:
			inputData = append(inputData, byte(r))
			fmt.Print(string(r))
		}
	}

	return string(inputData), nil
}

// validateInput 验证用户输入
func (q *PhoneNumber) validateInput(input string, def string) string {
	if input == "" {
		return def
	}
	return input
}

// maskPhoneNumber 对电话号码进行遮掩处理
func (q *PhoneNumber) maskPhoneNumber(number []byte) []byte {
	switch {
	case len(number) <= 3:
		return q.maskAll(number)
	case len(number) <= 7:
		return q.maskMiddle(number, 3)
	default:
		return q.preserveFirstThreeLastFive(number)
	}
}

// maskAll 将整个号码替换为星号
func (q *PhoneNumber) maskAll(number []byte) []byte {
	return bytes.Repeat([]byte{'*'}, len(number))
}

// maskMiddle 仅显示前几位数字，其余用星号代替
func (q *PhoneNumber) maskMiddle(number []byte, visibleDigits int) []byte {
	out := make([]byte, len(number))
	copy(out, number)
	for i := visibleDigits; i < len(number); i++ {
		out[i] = '*'
	}
	return out
}

// preserveFirstThreeLastFive 保留前三后五位数字，中间部分用星号代替
func (q *PhoneNumber) preserveFirstThreeLastFive(number []byte) []byte {
	out := make([]byte, len(number))
	copy(out, number)
	for i := 3; i < len(number)-4; i++ {
		out[i] = '*'
	}
	return out
}

// clearPreviousInput 清除之前的输入字符
func (q *PhoneNumber) clearPreviousInput(length int) {
	for i := 0; i < length; i++ {
		q.eraseLastChar()
	}
}

// eraseLastChar 删除最后一个字符并替换为空格
func (q *PhoneNumber) eraseLastChar() {
	fmt.Print("\b \b")
}

func (q *PhoneNumber) println(val string) {
	// 光标移动到行末
	fmt.Printf("%s\r\n", val)
}
