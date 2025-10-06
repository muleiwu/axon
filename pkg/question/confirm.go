package question

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"github.com/gookit/color"
)

type Confirm struct {
}

func NewConfirm() *Confirm {
	return &Confirm{}
}

func (q *Confirm) Execute(question string, def any) any {
	if strings.Contains(question, "</>") {
		color.Println(question)
	} else {
		color.Info.Println(question)
	}
	if def.(bool) {
		color.Info.Print("默认: y")
	} else {
		color.Info.Print("默认: n")
	}
	color.Info.Print(" [")
	color.LightBlue.Print("Y")
	color.Info.Print("es/")
	color.LightRed.Print("N")
	color.Info.Print("o] : ")

	input, _ := q.readInput()
	return q.validateInput(input, def.(bool))
}

func (q *Confirm) readInput() (string, error) {
	// 设置终端为原始模式，禁用缓冲
	oldState, err := q.setRawMode()
	if err != nil {
		return "", err
	}
	defer q.restoreMode(oldState)

	reader := bufio.NewReader(os.Stdin)
	var input string

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			return "", err
		}

		// 只接受YyNn字符
		switch char {
		case 'Y', 'y', 'N', 'n':
			// 如果有前一次输入，先删除它
			if input != "" {
				fmt.Print("\b \b") // 删除前一个字符
			}
			// 输出新字符
			fmt.Printf("%c", char)
			input = string(char)
		case '\r', '\n':
			// 按回车确认
			if input != "" {
				fmt.Println()
				return input, nil
			}
			// 如果还没有输入有效字符，忽略回车
			continue
		default:
			// 忽略其他字符，不显示
			continue
		}
	}
}

// setRawMode 设置终端为原始模式
func (q *Confirm) setRawMode() (*syscall.Termios, error) {
	var oldState syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(os.Stdin.Fd()),
		uintptr(syscall.TIOCGETA), uintptr(unsafe.Pointer(&oldState)), 0, 0, 0)
	if err != 0 {
		return nil, err
	}

	newState := oldState
	newState.Lflag &^= syscall.ICANON | syscall.ECHO
	newState.Cc[syscall.VMIN] = 1
	newState.Cc[syscall.VTIME] = 0

	_, _, err = syscall.Syscall6(syscall.SYS_IOCTL, uintptr(os.Stdin.Fd()),
		uintptr(syscall.TIOCSETA), uintptr(unsafe.Pointer(&newState)), 0, 0, 0)
	if err != 0 {
		return nil, err
	}

	return &oldState, nil
}

// restoreMode 恢复终端模式
func (q *Confirm) restoreMode(oldState *syscall.Termios) {
	syscall.Syscall6(syscall.SYS_IOCTL, uintptr(os.Stdin.Fd()),
		uintptr(syscall.TIOCSETA), uintptr(unsafe.Pointer(oldState)), 0, 0, 0)
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
