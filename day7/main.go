package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	textSlice := strings.Split(text, ",")
	crabs := make(Crabs, len(textSlice))
	for i := range textSlice {
		num, _ := strconv.Atoi(textSlice[i])
		crabs[i] = num
	}

	bestPos, fuel := crabs.calcBestPosition("simple")
	fmt.Println("answer part 1: ", bestPos)
	fmt.Println("fuel part 1: ", fuel)

	bestPos, fuel = crabs.calcBestPosition("complex")
	fmt.Println("answer part 2: ", bestPos)
	fmt.Println("fuel part 2: ", fuel)
}

type Crabs []int

func (c Crabs) calcBestPosition(s string) (int, int) {
	bestPos := math.MaxInt
	minFuel := math.MaxInt
	min, max := minMaxValue(c)
	for i := min; i < max; i++ {
		currentFuel := 0
		if s == "simple" {
			currentFuel = c.calcFuelToPosition(i)
		} else if s == "complex" {
			currentFuel = c.calcFuelToPositionComplex(i)
		} else {
			log.Fatal("type not suported")
		}
		if currentFuel < minFuel {
			bestPos = i
			minFuel = currentFuel
		}
	}
	return bestPos, minFuel
}

func (c Crabs) calcFuelToPosition(p int) int {
	fuel := 0
	for _, v := range c {
		fuel += abs(v - p)
	}
	return fuel
}

func (c Crabs) calcFuelToPositionComplex(p int) int {
	fuel := 0
	for _, v := range c {
		dist := abs(v - p)
		fuel += dist * (dist + 1) / 2
	}
	return fuel
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func minMaxValue(vals []int) (int, int) {
	min := math.MaxInt
	max := 0
	for _, v := range vals {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return min, max
}
