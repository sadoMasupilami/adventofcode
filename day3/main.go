package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const arraySize = 12
const listLength = 1000

func main() {
	bytes, err := ioutil.ReadFile("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)

	report := Report{
		input: text,
	}

	report.Process()
	fmt.Println("Solution part 1: ", report.gamma*report.epsilon)
	ogr := reduce(report.twoDimensional, 0, 1)
	csr := reduce(report.twoDimensional, 0, 0)
	fmt.Println("Solution part 2: ", ogr*csr)
}

type Report struct {
	input              string
	splitInput         []string
	twoDimensional     [][]int
	counter0, counter1 [arraySize]int
	gamma, epsilon     int64
}

func (r *Report) Process() {
	r.splitInput = strings.Split(r.input, "\n")

	r.twoDimensional = make([][]int, listLength)
	for i := range r.twoDimensional {
		r.twoDimensional[i] = make([]int, arraySize)
	}
	for x, s := range r.splitInput {
		for y, char := range s {
			if int(char) == 48 {
				r.twoDimensional[x][y] = 0
				r.counter0[y] += 1
			} else if int(char) == 49 {
				r.twoDimensional[x][y] = 1
				r.counter1[y] += 1
			} else {
				log.Fatal("your input is wrong, should only be 0 or 1 at every digit")
			}
		}
	}

	var sbg, sbe strings.Builder
	for pos, _ := range r.counter0 {
		if r.counter1[pos] >= r.counter0[pos] {
			sbg.WriteRune(49)
			sbe.WriteRune(48)
		} else {
			sbg.WriteRune(48)
			sbe.WriteRune(49)
		}
	}
	r.gamma, _ = strconv.ParseInt(sbg.String(), 2, 64)
	r.epsilon, _ = strconv.ParseInt(sbe.String(), 2, 64)
}

func reduce(input [][]int, digit int, dominantNumber int) int {
	var reducedArray [][]int

	target := countBinaries(input, digit, dominantNumber)
	for i, _ := range input {
		if input[i][digit] == target {
			reducedArray = append(reducedArray, input[i])
		}
	}
	fmt.Println("TEST: ", reducedArray)

	if len(reducedArray) <= 1 {
		sb := strings.Builder{}
		for _, val := range reducedArray[0] {
			sb.WriteString(strconv.Itoa(val))
		}
		number, err := strconv.ParseInt(sb.String(), 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		return int(number)
		//value, err := strconv.ParseInt()
	}
	return reduce(reducedArray, digit+1, dominantNumber)
}

func countBinaries(input [][]int, digit int, dominantNumber int) int {
	var zeroes, ones int
	for _, line := range input {
		if line[digit] == 1 {
			ones++
		} else {
			zeroes++
		}
	}
	var answer int
	fmt.Println("DEBUG 0: ", zeroes, " 1: ", ones)
	if zeroes > ones {
		answer = 0
	} else if ones > zeroes {
		answer = 1
	} else {
		if dominantNumber == 0 {
			return 0
		} else {
			return 1
		}
	}
	if dominantNumber == 1 {
		return answer
	} else {
		return 1 - answer
	}
}
