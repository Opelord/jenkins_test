package task4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bingo struct {
	Numbers []int
	Boards  []*Board
}

type Board struct {
	Numbers       [][]int
	Checked       [][]bool
	SumOfUnmarked int
}

func NewBoard() *Board {
	board := &Board{
		Numbers: make([][]int, 5),
		Checked: make([][]bool, 5),
	}
	for i := range board.Numbers {
		board.Numbers[i] = make([]int, 5)
		board.Checked[i] = make([]bool, 5)
	}
	return board
}

func (b *Board) CheckField(called int) bool {
	for i, row := range b.Numbers {
		for j, n := range row {
			if n != called {
				continue
			}
			b.Checked[i][j] = true
			b.SumOfUnmarked -= called
			for k := 0; k < 5; k++ {
				if !b.Checked[i][k] {
					break
				}
				if k == 4 {
					return true
				}
			}
			for k := 0; k < 5; k++ {
				if !b.Checked[k][j] {
					break
				}
				if k == 4 {
					return true
				}
			}
		}
	}

	return false
}

func ReadData(name string) (*Bingo, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("couldn't read file: %v", err)
	}
	defer f.Close()

	bingo := Bingo{
		Numbers: []int{},
		Boards:  []*Board{},
	}

	scanner := bufio.NewScanner(f)
	// read first line
	scanner.Scan()
	numString := scanner.Text()
	numSplit := strings.Split(numString, ",")
	for _, num := range numSplit {
		n, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		bingo.Numbers = append(bingo.Numbers, n)
	}

	var counter int
	for scanner.Scan() {
		if scanner.Text() == "" {
			board := NewBoard()
			bingo.Boards = append(bingo.Boards, board)
			continue
		}

		numSplit = strings.Fields(scanner.Text())
		for i, num := range numSplit {
			n, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			bingo.Boards[counter/5].Numbers[counter%5][i] = n
			bingo.Boards[counter/5].SumOfUnmarked += n
		}
		counter++
	}

	return &bingo, nil
}

func SolvePart1(name string) int {
	bingo, err := ReadData(name)
	if err != nil {
		log.Fatalf("couldn't read data: %v", err)
	}

	for _, n := range bingo.Numbers {
		for _, board := range bingo.Boards {
			if board.CheckField(n) {
				return board.SumOfUnmarked * n
			}
		}
	}

	return 0
}
