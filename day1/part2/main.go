package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1/part1/measurements.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var values []int
	scanner.Scan()
	firstValue, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	values = append(values, firstValue)
	scanner.Scan()
	secondValue, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	values[0] += secondValue
	values = append(values, secondValue)

	counter := 2
	for scanner.Scan() {
		currentValue, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		values[counter-2] += currentValue
		values[counter-1] += currentValue
		values = append(values, currentValue)
		counter++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	values = values[:len(values)-2]
	fmt.Println(values)

	increases := 0
	for i := 1; i < len(values); i++ {
		if values[i] > values[i-1] {
			increases++
		}
	}

	fmt.Println("increases: ", increases)
}
