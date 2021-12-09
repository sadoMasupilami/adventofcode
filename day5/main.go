package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const size = 1000

func main() {
	bytes, err := ioutil.ReadFile("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	textSlice := strings.Split(text, "\n")

	var vents []Vent
	for _, line := range textSlice {
		splittedLine := strings.Split(line, " -> ")
		vent := Vent{
			from: constructPoint(splittedLine[0]),
			to:   constructPoint(splittedLine[1]),
		}
		vents = append(vents, vent)
	}

	var field VentField
	field = make([][]int, size)
	for i := range field {
		field[i] = make([]int, size)
	}

	for _, vent := range vents {
		field.ApplyVent(vent, false)
	}
	fmt.Println("answer part 1: ", field.GetOverlaps())

	field.Reset()
	for _, vent := range vents {
		field.ApplyVent(vent, true)
	}
	fmt.Println("answer part 2: ", field.GetOverlaps())
}

type VentField [][]int

func (f VentField) ApplyVent(v Vent, withDiag bool) {
	if v.from.x == v.to.x {
		if v.from.y > v.to.y {
			v.from.y, v.to.y = v.to.y, v.from.y
		}
		for i := v.from.y; i < v.to.y+1; i++ {
			p := Point{
				x: v.from.x,
				y: i,
			}
			f.ApplyPoint(p)
		}

	} else if v.from.y == v.to.y {
		if v.from.x > v.to.x {
			v.from.x, v.to.x = v.to.x, v.from.x
		}
		for i := v.from.x; i < v.to.x+1; i++ {
			p := Point{
				x: i,
				y: v.from.y,
			}
			f.ApplyPoint(p)
		}

	} else if withDiag {
		if v.from.x > v.to.x && v.from.y > v.to.y || v.from.x < v.to.x && v.from.y > v.to.y {
			v.from.x, v.to.x = v.to.x, v.from.x
			v.from.y, v.to.y = v.to.y, v.from.y
		}
		if v.from.x < v.to.x && v.from.y < v.to.y {
			for i := 0; i <= v.to.x-v.from.x; i++ {
				p := Point{
					x: v.from.x + i,
					y: v.from.y + i,
				}
				f.ApplyPoint(p)
			}
		} else if v.from.x > v.to.x && v.from.y < v.to.y {
			for i := 0; i <= v.from.x-v.to.x; i++ {
				p := Point{
					x: v.from.x - i,
					y: v.from.y + i,
				}
				f.ApplyPoint(p)
			}
		}
	}
}

func (f VentField) ApplyPoint(p Point) {
	f[p.y][p.x] += 1
}

func (f VentField) GetOverlaps() int {
	overlaps := 0
	for _, line := range f {
		for _, val := range line {
			if val > 1 {
				overlaps++
			}
		}
	}
	return overlaps
}

func (f VentField) Reset() {
	for i, line := range f {
		for j := range line {
			f[i][j] = 0
		}
	}
}

func (f VentField) String() string {
	sb := strings.Builder{}
	for _, v := range f {
		sb.WriteString(fmt.Sprintf("%v\n", v))
	}
	return sb.String()
}

type Vent struct {
	from Point
	to   Point
}

func (v Vent) String() string {
	return fmt.Sprintf("FROM: %v\nTO: %v\n", v.from, v.to)
}

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("X: %v, Y: %v", p.x, p.y)
}

func constructPoint(s string) Point {
	coords := strings.Split(s, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	return Point{x, y}
}
