package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	textSlice := strings.Split(text, ",")

	var fish1 FishSwarm
	fish2 := make(FishSwarm, 9)
	for _, str := range textSlice {
		num, _ := strconv.Atoi(str)
		fish1 = append(fish1, num)
		fish2[num]++
	}

	for i := 0; i < 18; i++ {
		fish1 = fish1.Reproduce()
	}
	fmt.Println("Solution part 1: ", len(fish1))
	fmt.Println("Solution part 2: ", fish2.ReproduceFast(256))
}

type FishSwarm []int

func (s *FishSwarm) ReproduceFast(count int) int {
	for i := 0; i < count; i++ {
		newFish := (*s)[0]
		(*s)[0] = (*s)[1]
		(*s)[1] = (*s)[2]
		(*s)[2] = (*s)[3]
		(*s)[3] = (*s)[4]
		(*s)[4] = (*s)[5]
		(*s)[5] = (*s)[6]
		(*s)[6] = (*s)[7]
		(*s)[6] += newFish
		(*s)[7] = (*s)[8]
		(*s)[8] = newFish
	}
	fish := 0
	for i := range *s {
		fish += (*s)[i]
	}
	return fish
}

func (s FishSwarm) Reproduce() FishSwarm {
	ns := FishSwarm{}
	newFishCounter := 0
	for i, f := range s {
		s[i] = f - 1
	}
	for _, f := range s {
		if f < 0 {
			f = 6
			newFishCounter++
		}
		ns = append(ns, f)
	}
	for i := 0; i < newFishCounter; i++ {
		ns = append(ns, 8)
	}
	return ns
}
