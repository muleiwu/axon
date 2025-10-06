package question

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gookit/color"
	"golang.org/x/term"
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
	color.LightBlue.Print("Y")
	color.Info.Print("es/")
	color.LightRed.Print("N")
	color.Info.Print("o] : ")

	input, _ := q.readInput(def.(bool))
	return q.validateInput(input, def.(bool))
}

func (q *Confirm) readInput(def bool) (string, error) {
	// 设置终端为原始模式，禁用缓冲
	oldState, err := q.setRawMode()
	if err != nil {
		return "", err
	}
	defer q.restoreMode(oldState)

	// 设置信号处理，确保Ctrl+C能正常终止程序
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigChan)

	reader := bufio.NewReader(os.Stdin)
	var input string

	// 显示默认值
	if def {
		fmt.Print("Y")
		input = "Y"
	} else {
		fmt.Print("N")
		input = "N"
	}

	for {
		// 使用select来同时监听键盘输入和信号
		select {
		case <-sigChan:
			// 收到中断信号，恢复终端模式并退出
			q.restoreMode(oldState)
			fmt.Println() // 换行
			os.Exit(0)
		default:
			// 非阻塞读取键盘输入
		}

		char, _, err := reader.ReadRune()
		if err != nil {
			return "", err
		}

		// 处理Ctrl+C (ASCII 3)
		if char == 3 {
			q.restoreMode(oldState)
			fmt.Println() // 换行
			os.Exit(0)
		}

		// 处理删除键 (Backspace = 127, Delete = 8)
		if char == 127 || char == 8 {
			if input != "" {
				fmt.Print("\b \b") // 删除前一个字符
				input = ""
			}
			continue
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
			fmt.Println()
			return input, nil
		default:
			// 忽略其他字符，不显示
			continue
		}
	}
}

// setRawMode 设置终端为原始模式（跨平台兼容）
func (q *Confirm) setRawMode() (interface{}, error) {
	// 使用golang.org/x/term包，自动处理不同平台的差异
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}
	return oldState, nil
}

// restoreMode 恢复终端模式（跨平台兼容）
func (q *Confirm) restoreMode(oldState interface{}) {
	if oldState != nil {
		if state, ok := oldState.(*term.State); ok {
			term.Restore(int(os.Stdin.Fd()), state)
		}
	}
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
