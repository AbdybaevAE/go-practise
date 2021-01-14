package main

import (
	"fmt"
	"strings"
)
type Quiz struct {
	questions []Question
	inited bool
}
type Question struct {
	firstNum string
	secondNum string
	answer string
	number int
}

func (question *Question) GetQuestion () string {
	return fmt.Sprintf("Question #%d: How much would be %s + %s?", question.number, question.firstNum, question.secondNum)
}
func (quiz *Quiz) addQuestion(content string, answer string) {
	nums := strings.Split(content, "+")
	if len(nums) != 2 {
		panic("Wrong question!")
	}
	question := Question{nums[0], nums[1], answer, len(quiz.questions ) + 1}
	quiz.questions = append(quiz.questions, question)
}
func (quiz *Quiz) init() {
	if quiz.inited {
		fmt.Println("already inited...")
		return
	}
	fmt.Println("initialize...")
	quiz.questions = make([] Question, 0)
	quiz.inited = true

}
