package question

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

// Selection 选择问题结构体
type Selection struct{}

func NewSelection() *Selection {
	return &Selection{}
}

// SelectionItem 选择项
type SelectionItem struct {
	Label       string // 显示的标签
	Description string // 描述信息
	Value       string // 实际值
}

// Execute 执行选择问题
// question: 问题提示
// items: 选择项列表
// 返回: 用户选择的项目的Value
func (s *Selection) Execute(question string, items []SelectionItem) string {
	// 显示问题标题
	color.Cyan.Printf("\n%s\n", question)
	color.Gray.Println(strings.Repeat("-", len(question)))

	// 显示选择项
	for i, item := range items {
		color.Printf("  <cyan>%d</>. <green>%s</> - %s\n", i+1, item.Label, item.Description)
	}

	// 显示退出选项
	color.Printf("  <cyan>%d</>. <red>退出</>\n", len(items)+1)

	// 获取用户输入
	for {
		color.Print("\n请选择 (输入数字): ")

		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			color.Red.Println("输入错误，请重新输入")
			continue
		}

		// 验证输入
		choice, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			color.Red.Println("请输入有效的数字")
			continue
		}

		// 退出选项
		if choice == len(items)+1 {
			return "exit"
		}

		// 验证选择范围
		if choice < 1 || choice > len(items) {
			color.Red.Printf("请输入 1-%d 之间的数字\n", len(items)+1)
			continue
		}

		// 返回选择的值
		return items[choice-1].Value
	}
}
