package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	textSlice := strings.Split(text, "\n")

	numbersRaw := strings.Split(textSlice[0], ",")
	var numbers []int
	for _, val := range numbersRaw {
		num, _ := strconv.Atoi(val)
		numbers = append(numbers, num)
	}

	var loserBoards []BingoBoard
	var winnerBoards []BingoBoard
	for i := 2; i < len(textSlice); i += 6 {
		loserBoard := BingoBoard{}
		winnerBoard := BingoBoard{}
		for i, val := range textSlice[i : i+5] {
			line := strings.Fields(strings.TrimSpace(val))
			for j, num := range line {
				realNum, _ := strconv.Atoi(num)
				winnerBoard.Numbers[i][j] = realNum
				loserBoard.Numbers[i][j] = realNum
			}
		}
		winnerBoards = append(winnerBoards, winnerBoard)
		loserBoards = append(loserBoards, loserBoard)
	}

	fmt.Println("answer part 1: ", GetWinnerValue(numbers, winnerBoards))
	fmt.Println("answer part 2: ", GetLoserValue(numbers, loserBoards))

}

func GetWinnerValue(numbers []int, boards []BingoBoard) int {
	for _, num := range numbers {
		for i := range boards {
			boards[i].CallNumber(num)
			bingo, answer := boards[i].IsBingo(num)
			if bingo {
				return answer
			}
		}
	}
	return 0
}

func GetLoserValue(numbers []int, boards []BingoBoard) int {
	numberOfBoards := len(boards)
	numberOfWinnerBoards := 0
	for _, num := range numbers {
		for i, b := range boards {
			alreadyBingo, _ := b.IsBingo(0)
			if alreadyBingo {
				continue
			}
			boards[i].CallNumber(num)
			bingo, answer := boards[i].IsBingo(num)
			if bingo {
				numberOfWinnerBoards++
				if numberOfWinnerBoards >= numberOfBoards {
					return answer
				}
			}
		}
	}
	return 0
}

type BingoBoard struct {
	Numbers [5][5]int
	Called  [5][5]bool
}

func (b BingoBoard) String() string {
	sb := strings.Builder{}
	sb.WriteString("-----------------------------\n")
	for i, line := range b.Numbers {
		for j, num := range line {
			numAsString := strconv.Itoa(num)
			sb.WriteString("  " + numAsString)
			if b.Called[i][j] {
				sb.WriteString("X")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (b BingoBoard) IsBingo(lastNumber int) (bool, int) {
	for i := 0; i < 5; i++ {
		if b.Called[i][0] && b.Called[i][1] && b.Called[i][2] && b.Called[i][3] && b.Called[i][4] {
			return true, lastNumber * b.SumUncalledNumbers()
		} else if b.Called[0][i] && b.Called[1][i] && b.Called[2][i] && b.Called[3][i] && b.Called[4][i] {
			return true, lastNumber * b.SumUncalledNumbers()
		}
	}
	return false, 0
}

func (b BingoBoard) SumUncalledNumbers() int {
	sum := 0
	for i, line := range b.Numbers {
		for j, num := range line {
			if !b.Called[i][j] {
				sum += num
			}
		}
	}
	return sum
}

func (b *BingoBoard) CallNumber(number int) {
	for i, line := range b.Numbers {
		for j := range line {
			if b.Numbers[i][j] == number {
				b.Called[i][j] = true
			}
		}
	}
}
