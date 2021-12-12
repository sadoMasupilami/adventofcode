package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	textSlice := strings.Split(text, "\n")

	var segmentRows []segmentRow
	for _, line := range textSlice {
		tempSplit := strings.Split(line, "|")
		patternsAll, outputAll := tempSplit[0], tempSplit[1]
		patterns := strings.Split(patternsAll, " ")[:10]
		output := strings.Split(outputAll, " ")[1:]
		for i := range patterns {
			patterns[i] = orderString(patterns[i])
		}
		for i := range output {
			output[i] = orderString(output[i])
		}
		sr := segmentRow{
			patterns: patterns,
			output:   output,
		}
		segmentRows = append(segmentRows, sr)
	}

	answer1 := 0
	for _, row := range segmentRows {
		answer1 += row.calc1()
	}
	fmt.Println("answer 1 : ", answer1)

	answer2 := 0
	for _, row := range segmentRows {
		answer2 += row.calc2()
	}
	fmt.Println("answer 2 : ", answer2)

}

type segmentRow struct {
	patterns []string
	output   []string
}

func (r segmentRow) calc1() int {
	val := 0
	for _, o := range r.output {
		if len(o) == 2 || len(o) == 4 || len(o) == 3 || len(o) == 7 {
			val++
		}
	}
	return val
}

func (r segmentRow) calc2() int {
	mappingStringToValue := make(map[string]int, 10)
	mappingValueToString := make(map[int]string, 10)
	for _, val := range r.patterns {
		if len(val) == 2 {
			mappingStringToValue[val] = 1
			mappingValueToString[1] = val
		}
	}
	for _, val := range r.patterns {
		if len(val) == 3 {
			mappingStringToValue[val] = 7
			mappingValueToString[7] = val
		}
	}
	for _, val := range r.patterns {
		if len(val) == 4 {
			mappingStringToValue[val] = 4
			mappingValueToString[4] = val
		}
	}
	for _, val := range r.patterns {
		if len(val) == 7 {
			mappingStringToValue[val] = 8
			mappingValueToString[8] = val
		}
	}
	for _, val := range r.patterns {
		if len(val) == 5 {
			if containsAllLetters(val, mappingValueToString[1]) {
				mappingStringToValue[val] = 3
				mappingValueToString[3] = val
			}
		}
	}
	for _, val := range r.patterns {
		if len(val) == 6 {
			if containsAllLetters(val, mappingValueToString[4]) {
				mappingStringToValue[val] = 9
				mappingValueToString[9] = val
			}
		}
	}
	for _, val := range r.patterns {
		if len(val) == 6 {
			if !containsAllLetters(val, mappingValueToString[1]) {
				mappingStringToValue[val] = 6
				mappingValueToString[6] = val
			}
		}
	}
	for _, val := range r.patterns {
		if len(val) == 5 {
			if containsAllLetters(mappingValueToString[6], val) {
				mappingStringToValue[val] = 5
				mappingValueToString[5] = val
			}
		}
	}
	for _, val := range r.patterns {
		if len(val) == 5 {
			if !containsAllLetters(val, mappingValueToString[3]) && !containsAllLetters(val, mappingValueToString[5]) {
				mappingStringToValue[val] = 2
				mappingValueToString[2] = val
			}
		}
	}

	digits := make([]int, 4)
	for i, val := range r.output {
		digits[i] = mappingStringToValue[val]
	}
	return digits[0]*1000 + digits[1]*100 + digits[2]*10 + digits[3]
}

func orderString(s string) string {
	rs := []rune(s)
	for i := range rs {
		for j := range rs {
			if rs[i] < rs[j] {
				rs[i], rs[j] = rs[j], rs[i]
			}
		}
	}
	return string(rs)
}

func containsAllLetters(s string, search string) bool {
	for _, val := range search {
		if !strings.Contains(s, string(val)) {
			return false
		}
	}
	return true
}
