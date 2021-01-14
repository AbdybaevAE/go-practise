package main

type AnswerSheet struct {
	answers      []Answer
	quiz         Quiz
	isStarted    bool
	isEnded      bool
	currQuestion int
	score        int
}

func (ans *AnswerSheet) receiveAnswer(clientAns string) bool {
	if ans.isEnded {
		return false
	}
	answer := &ans.answers[ans.currQuestion]
	answer.hasAnswer = true
	answer.content = clientAns
	ans.currQuestion++
	if ans.currQuestion == len(ans.answers)-1 {
		ans.endQuiz()
	}
	return true
}
func (ans *AnswerSheet) getQuestion() string {
	if ans.isEnded {
		return "Quiz has already ended..."
	}
	return ans.quiz.questions[ans.currQuestion].GetQuestion()

}
func (ans *AnswerSheet) endQuiz() {
	if ans.isEnded {
		return
	}
	ans.isEnded = true
	ans.computeResults()

}
func (ans *AnswerSheet) computeResults() {
	score := 0
	answers := ans.answers
	questions := ans.quiz.questions
	for i := range answers {
		if answers[i].hasAnswer && answers[i].content == questions[i].answer {
			score++
		}
	}
	ans.score = score
}
func startQuiz(quiz Quiz) AnswerSheet {
	count := len(quiz.questions)
	ans := AnswerSheet{answers: make([]Answer, count), quiz: quiz, isStarted: true, isEnded: false, currQuestion: 0}
	for i := range ans.answers {
		ans.answers[i] = Answer{content: "", hasAnswer: false}
	}
	return ans
}

type Answer struct {
	content   string
	hasAnswer bool
}
