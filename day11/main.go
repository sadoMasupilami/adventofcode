package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	textSlice := strings.Split(text, "\n")

	var numbers [][]int
	numbers = make([][]int, len(textSlice))
	for i, line := range textSlice {
		numbers[i] = make([]int, len(textSlice))
		for j, char := range line {
			num, _ := strconv.Atoi(string(char))
			numbers[i][j] = num
		}
	}

	field := Field{size: len(numbers)}
	field.Initialize(numbers)
	field.Cycles(100)
	fmt.Println("FLASHES: ", field.numFlashes)

	field.Initialize(numbers)
	answer2 := field.findSimultaneousCycle()
	fmt.Println("ALL FLASHED AT: ", answer2+1)
}

type Octopus struct {
	charge    int
	neighbors []*Octopus
	flashed   bool
}

type Field struct {
	octopuses  [][]Octopus
	size       int
	numFlashes int
}

func (f *Field) Initialize(numbers [][]int) {
	f.octopuses = make([][]Octopus, f.size)
	for i := range numbers {
		f.octopuses[i] = make([]Octopus, f.size)
		for j := range numbers[i] {
			f.octopuses[i][j] = Octopus{
				charge:    numbers[i][j],
				neighbors: nil,
				flashed:   false,
			}
		}
	}

	for i := range numbers {
		for j := range numbers[i] {
			for k := i - 1; k <= i+1; k++ {
				for l := j - 1; l <= j+1; l++ {
					if !(i == k && j == l) && k >= 0 && k < f.size && l >= 0 && l < f.size {
						f.octopuses[i][j].neighbors = append(f.octopuses[i][j].neighbors, &f.octopuses[k][l])
					}
				}
			}
		}
	}
}

func (o *Octopus) Increase(f *Field) {
	o.charge++
	if o.charge > 9 {
		o.Flash(f)
	}
}

func (o *Octopus) Flash(f *Field) {
	if o.flashed {
		return
	}
	o.flashed = true
	for _, neighbor := range o.neighbors {
		neighbor.Increase(f)
	}
	f.numFlashes++
}

func (f *Field) Cycles(c int) {
	for i := 0; i < c; i++ {
		f.Cycle()
	}
}

func (f *Field) Cycle() {
	for i := range f.octopuses {
		for j := range f.octopuses[i] {
			f.octopuses[i][j].Increase(f)
		}
	}
	f.EndRound()
}

func (f *Field) EndRound() {
	for i := range f.octopuses {
		for j := range f.octopuses[i] {
			if f.octopuses[i][j].flashed {
				f.octopuses[i][j].charge = 0
				f.octopuses[i][j].flashed = false
			}
		}
	}
}

func (f *Field) findSimultaneousCycle() int {
	for i := 0; i < 9999; i++ {
		f.Cycle()
		if f.haveAllFlashed() {
			return i
		}
	}
	return -1
}

func (f Field) haveAllFlashed() bool {
	for i := range f.octopuses {
		for j := range f.octopuses[i] {
			if f.octopuses[i][j].charge != 0 {
				return false
			}
		}
	}
	return true
}

func (f Field) String() string {
	sb := strings.Builder{}
	for i := range f.octopuses {
		for j := range f.octopuses[i] {
			sb.WriteString(fmt.Sprint(f.octopuses[i][j].charge))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
