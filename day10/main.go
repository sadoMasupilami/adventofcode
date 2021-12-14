package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	textSlice := strings.Split(text, "\n")

	highScore := 0
	var lineCompletionScores []int
	pointTranslation := map[int]int{
		3:     1,
		57:    2,
		1197:  3,
		25137: 4,
	}

	for _, line := range textSlice {
		s := stack{}
		var score bool
	Outer:
		for _, char := range line {
			score = true
			c := string(char)
			switch c {
			case "(":
				s = s.push(3)
			case ")":
				tempS, val := s.pop()
				if val != 3 {
					highScore += 3
					score = false
					break Outer
				} else {
					s = tempS
				}
			case "[":
				s = s.push(57)
			case "]":
				tempS, val := s.pop()
				if val != 57 {
					highScore += 57
					score = false
					break Outer
				} else {
					s = tempS
				}
			case "{":
				s = s.push(1197)
			case "}":
				tempS, val := s.pop()
				if val != 1197 {
					highScore += 1197
					score = false
					break Outer
				} else {
					s = tempS
				}
			case "<":
				s = s.push(25137)
			case ">":
				tempS, val := s.pop()
				if val != 25137 {
					highScore += 25137
					score = false
					break Outer
				} else {
					s = tempS
				}
			}
		}
		if score {
			lineScore := 0
			for _ = range s {
				tempS, val := s.pop()
				s = tempS
				lineScore = lineScore*5 + pointTranslation[val]
			}
			lineCompletionScores = append(lineCompletionScores, lineScore)
		}
	}
	fmt.Println("answer part 1: ", highScore)
	sort.Ints(lineCompletionScores)
	fmt.Println("answer part 2: ", lineCompletionScores[len(lineCompletionScores)/2])
}

type stack []int

func (s stack) push(v int) stack {
	return append(s, v)
}

func (s stack) pop() (stack, int) {
	return s[:len(s)-1], s[len(s)-1]
}
