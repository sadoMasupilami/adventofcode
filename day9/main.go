package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	textSlice := strings.Split(text, "\n")

	c := caves{}
	c = make([][]int, len(textSlice)+2)
	for i := range c {
		c[i] = make([]int, len(textSlice[0])+2)
	}

	numbers := make([][]int, len(textSlice))
	for i := range numbers {
		numbers[i] = make([]int, len(textSlice[0]))
	}
	for i, line := range textSlice {
		for j, val := range line {
			num, _ := strconv.Atoi(string(val))
			numbers[i][j] = num
		}
	}

	for i := range c {
		for j := range c[i] {
			if i == 0 || i == len(c)-1 || j == 0 || j == len(c[0])-1 {
				c[i][j] = 9
			} else {
				c[i][j] = numbers[i-1][j-1]
			}
		}
	}

	lowPoints := 0
	for i, line := range c[1 : len(c)-1] {
		for j, val := range line[1 : len(line)-1] {
			if val < c[i][j+1] && val < c[i+1][j] && val < c[i+2][j+1] && val < c[i+1][j+2] {
				lowPoints += val + 1
			}
		}
	}

	fmt.Println("answer part 1: ", lowPoints)
}

type caves [][]int
