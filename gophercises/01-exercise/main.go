package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

const defaultFile = "/Users/azamatabdybaev/repos/play/go/practise/gophercises/01-exercise/data.csv"

func main() {
	filePath := flag.String("file", defaultFile, "File to import")
	flag.Parse()
	quiz := ParseCSV(*filePath)
	quiz.init()
	ans := startQuiz(quiz)
	func() {
		scanner := bufio.NewScanner(os.Stdin)
		for !ans.isEnded {
			title := ans.getQuestion()
			fmt.Println(title)
			scanner.Scan()
			currAns := scanner.Text()
			fmt.Printf("Accepted answer: %s \n", currAns)
			ans.receiveAnswer(currAns)
		}
	}()

	fmt.Println("Quiz was ended!")
	fmt.Printf("Your score: %d\n", ans.score)
	fmt.Println("Bye!")


}