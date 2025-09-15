package Interfaces

type QuestionInterface interface {
	Execute(question string, def any) any
}
