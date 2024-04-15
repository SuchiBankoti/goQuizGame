package main

import (
	"fmt"
	"os"
	"strings"
)

type Question struct {
	question string
	ans      string
}

func main() {
	myMap := make(map[int]Question)
	content, err := os.ReadFile("./text")
	if err != nil {
		fmt.Print(err)
		return
	}

	lines := strings.Split(string(content), ",")
	for index, line := range lines {
		temp := strings.Split(line, "=")
		question := temp[0]
		ans := temp[1]
		myMap[index] = Question{question: question, ans: ans}
	}
	for _, ques := range myMap {
		fmt.Printf("Question : %v:\n", ques.question)
		var answer string
		fmt.Scanln(&answer)

		if answer == ques.ans {
			fmt.Print("correct\n")
		} else {
			fmt.Print("incorrect\n")
		}
	}

}
