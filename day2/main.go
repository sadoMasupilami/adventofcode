package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	textSlice := strings.Split(text, "\n")

	// part 1
	forward, depth := 0, 0
	for _, s := range textSlice {
		direction := strings.Split(s, " ")[0]
		lengthAsString := strings.Split(s, " ")[1]
		length, err := strconv.Atoi(lengthAsString)
		if err != nil {
			log.Fatal(err)
		}
		switch direction {
		case "forward":
			forward += length
		case "down":
			depth += length
		case "up":
			depth -= length
		}
	}

	log.Println("forward: ", forward)
	log.Println("depth: ", depth)
	log.Println("answer: ", forward*depth)

	// part 2
	forward, depth = 0, 0
	aim := 0
	for _, s := range textSlice {
		direction := strings.Split(s, " ")[0]
		lengthAsString := strings.Split(s, " ")[1]
		length, err := strconv.Atoi(lengthAsString)
		if err != nil {
			log.Fatal(err)
		}
		switch direction {
		case "forward":
			forward += length
			depth += aim * length
		case "down":
			aim += length
		case "up":
			aim -= length
		}
		//log.Println("-----------------")
		//log.Println("STEP: ", s)
		//log.Println("aim: ", aim)
		//log.Println("forward: ", forward)
		//log.Println("depth: ", depth)
	}

	log.Println("answer: ", forward*depth)
}
