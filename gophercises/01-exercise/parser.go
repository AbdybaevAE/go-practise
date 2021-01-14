package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func ParseCSV(filePath string) Quiz {
	fmt.Println("Reading file...")
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(f)
	quiz := Quiz{}
	quiz.init()
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		quiz.addQuestion(record[0], record[1])
	}
	fmt.Println("End of reading...")
	return quiz
}
