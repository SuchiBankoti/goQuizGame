package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Question struct {
	question string
	ans      int
}

func main() {
	myMap := make(map[int]Question)
	var score int = 0
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
		ansNum, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Print(err)
			return
		}
		myMap[index] = Question{question: question, ans: ansNum}
	}
	for _, ques := range myMap {
		fmt.Printf("Question : %v:\n", ques.question)
		var answer int
		fmt.Scanln(&answer)

		if answer == ques.ans {
			score += 1
			fmt.Printf("correct\n score:%v/%v", score, len(myMap))
		} else {
			fmt.Printf("incorrect\n score:%v/%v", score, len(myMap))
		}
	}

}
