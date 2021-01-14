package main

import (
	"bufio"
	"fmt"
	"os"
)

const filePath = "/Users/azamatabdybaev/repos/play/go/practise/gophercises/01-exercise/data.csv"

func main() {
	quiz := ParseCSV(filePath)
	quiz.init()
	ans := startQuiz(quiz)
	scanner := bufio.NewScanner(os.Stdin)
	for !ans.isEnded {
		title := ans.getQuestion()
		fmt.Println(title)
		scanner.Scan()
		currAns := scanner.Text()
		fmt.Printf("Accepted answer: %s \n", currAns)
		ans.receiveAnswer(currAns)
	}
	fmt.Println("Quiz was ended!")
	fmt.Printf("Your score: %d\n", ans.score)
	fmt.Println("Bye!")


}