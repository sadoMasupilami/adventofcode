package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1/measurements.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := -1
	lastValue := 0
	currentValue := 0
	for scanner.Scan() {
		lastValue = currentValue
		currentValue, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if currentValue > lastValue {
			counter++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of increases: ", counter)
}
