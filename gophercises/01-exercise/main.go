package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

const defaultFile = "data.csv"

func main() {
	filePath := flag.String("file", defaultFile, "File to import")
	limit := flag.Int("limit", 2, "Time for quiz")
	inpCh := make(chan string)
	flag.Parse()
	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	fmt.Printf("Limit is %d\n", *limit)
	quiz := ParseCSV(*filePath)
	quiz.init()
	ans := startQuiz(quiz)
	scanner := bufio.NewScanner(os.Stdin)
actLabel:
	for !ans.isEnded {
		title := ans.getQuestion()
		fmt.Println(title)
		go func() {
			scanner.Scan()
			inpCh <- scanner.Text()
		}()
		select {
		case <-timer.C:
			break actLabel
		case currAns := <-inpCh:
			fmt.Printf("Accepted answer: %s \n", currAns)
			ans.receiveAnswer(currAns)
		}

	}
	ans.endQuiz()
	fmt.Println("Quiz was ended!")
	fmt.Printf("Your score: %d\n", ans.score)
	fmt.Println("Bye!")

}
