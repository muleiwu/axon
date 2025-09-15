# Axon

Axon 是一个用于构建交互式命令行应用的轻量级 Go 库，提供了简洁易用的终端交互功能。

## 功能特点

- 简单易用的 API，专注于提供优质的终端交互体验
- 支持多种交互方式：普通问答、确认选择、密码输入、电话号码输入和列表选择
- 内置数据验证和格式化功能
- 支持默认值设置
- 电话号码自动脱敏处理，保护用户隐私
- 彩色输出支持，增强视觉体验

## 安装

```bash
go get github.com/muleiwu/axon
```

## 快速开始

```go
package main

import (
	"fmt"
	"github.com/muleiwu/axon"
	"github.com/muleiwu/axon/pkg/question"
)

func main() {
	// 普通问题
	name := axon.Question("请输入您的姓名: ", "")
	fmt.Printf("您好，%s!\n", name)
	
	// 确认问题
	proceed := axon.Confirm("是否继续? (y/n) ", true)
	if proceed {
		fmt.Println("继续执行...")
	} else {
		fmt.Println("已取消")
		return
	}
	
	// 密码输入（隐藏显示）
	password := axon.Password("请输入密码: ", "")
	fmt.Println("密码已保存")
	
	// 电话号码（带验证和脱敏）
	phone := axon.PhoneNumber("请输入手机号: ", "")
	fmt.Println("手机号已保存")
	
	// 选择列表
	items := []question.SelectionItem{
		{Label: "选项一", Description: "这是第一个选项", Value: "option1"},
		{Label: "选项二", Description: "这是第二个选项", Value: "option2"},
		{Label: "选项三", Description: "这是第三个选项", Value: "option3"},
	}
	selected := axon.Selection("请选择一个选项:", items)
	
	if selected == "exit" {
		fmt.Println("您选择了退出")
	} else {
		fmt.Printf("您选择了: %s\n", selected)
	}
}
```

## API 参考

### 普通问题

```go
func Question(question string, defaultValue string) string
```

询问用户普通问题并获取文本输入。

- `question`: 要显示给用户的问题文本
- `defaultValue`: 默认值，当用户直接按回车时使用
- 返回: 用户输入的字符串

### 确认问题

```go
func Confirm(question string, defaultValue bool) bool
```

询问用户确认问题，获取 yes/no 回答。

- `question`: 要显示给用户的问题文本
- `defaultValue`: 默认值，当用户直接按回车时使用
- 返回: 用户的选择结果(true/false)

### 密码输入

```go
func Password(question string, defaultValue string) string
```

询问用户密码，输入过程中密码会被隐藏显示。

- `question`: 要显示给用户的问题文本
- `defaultValue`: 默认值，当用户直接按回车时使用
- 返回: 用户输入的密码字符串

### 电话号码输入

```go
func PhoneNumber(question string, defaultValue string) string
```

询问用户电话号码，包含格式验证功能。输入后自动脱敏处理，保护用户隐私。

- `question`: 要显示给用户的问题文本
- `defaultValue`: 默认值，当用户直接按回车时使用
- 返回: 经过验证的电话号码字符串

### 选择列表

```go
func Selection(question string, items []SelectionItem) string
```

让用户从列表中选择一项。

- `question`: 要显示给用户的问题文本
- `items`: 选择项列表
- 返回: 用户选择的项目的Value，如果选择退出则返回"exit"

## 自定义选择项

创建选择列表时，可以使用 `SelectionItem` 结构定义选项：

```go
type SelectionItem struct {
	Label       string // 显示的标签
	Description string // 描述信息
	Value       string // 实际值
}
```

## 许可证

本项目采用 [LICENSE](LICENSE) 许可证开源。
