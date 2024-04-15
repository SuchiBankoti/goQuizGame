package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Question struct {
	question string
	ans      int
}

var score int = 0
var total int

func main() {
	fmt.Print("set timer:")
	var limit int
	fmt.Scanln(&limit)
	go quiz()
	time.Sleep(time.Duration(limit) * time.Second)
	fmt.Printf("\nTime limit ended, your score is %v/%v", score, total)

}
func quiz() {
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
		ansNum, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Print(err)
			return
		}
		myMap[index] = Question{question: question, ans: ansNum}
	}
	total = len(myMap)
	for _, ques := range myMap {
		fmt.Printf("Problem: %v=", ques.question)
		var answer int
		fmt.Scanln(&answer)

		if answer == ques.ans {
			score += 1
			fmt.Printf("correct\n")
		} else {
			fmt.Printf("incorrect\n")
		}
	}
}
