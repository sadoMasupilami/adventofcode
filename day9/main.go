package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("day9/input-small.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	textSlice := strings.Split(text, "\n")

	fmt.Println(textSlice)

	c := caves{}
	c = make([][]int, len(textSlice)+2)
	for _, line := range c {
		line = make([]int, len(line)+2)
	}

}

type caves [][]int
