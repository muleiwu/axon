package axon

import question2 "github.com/muleiwu/axon/pkg/question"

// Confirm 询问用户确认问题
// question: 要显示给用户的问题文本
// defaultValue: 默认值，当用户直接按回车时使用
// 返回: 用户的选择结果(true/false)
func Confirm(question string, defaultValue bool) bool {
	confirm := question2.NewConfirm()
	return confirm.Execute(question, defaultValue).(bool)
}

// Question 询问用户普通问题
// question: 要显示给用户的问题文本
// defaultValue: 默认值，当用户直接按回车时使用
// 返回: 用户输入的字符串
func Question(question string, defaultValue string) string {
	newQuestion := question2.NewQuestion()
	return newQuestion.Execute(question, defaultValue).(string)
}

// PhoneNumber 询问用户电话号码
// 包含格式验证功能
// question: 要显示给用户的问题文本
// defaultValue: 默认值，当用户直接按回车时使用
// 返回: 经过验证的电话号码字符串
func PhoneNumber(question string, defaultValue string) string {
	phoneNumber := question2.NewPhoneNumber()
	return phoneNumber.Execute(question, defaultValue).(string)
}

// Password 询问用户密码
// 输入过程中密码会被隐藏显示
// question: 要显示给用户的问题文本
// defaultValue: 默认值，当用户直接按回车时使用
// 返回: 用户输入的密码字符串
func Password(question string, defaultValue string) string {
	password := question2.NewPassword()
	return password.Execute(question, defaultValue).(string)
}

// Selection 让用户从列表中选择一项
// question: 要显示给用户的问题文本
// items: 选择项列表
// 返回: 用户选择的项目的Value，如果选择退出则返回"exit"
func Selection(question string, items []question2.SelectionItem) string {
	selection := question2.NewSelection()
	return selection.Execute(question, items)
}
