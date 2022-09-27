package task4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadData(t *testing.T) {
	assert := assert.New(t)
	bingo, err := ReadData("../example_data")
	assert.NoError(err)

	bingoNums := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	assert.ElementsMatch(bingoNums, bingo.Numbers, "numbers in bingo doesn't match")
	secondBoard := [][]int{
		{3, 15, 0, 2, 22},
		{9, 18, 13, 17, 5},
		{19, 8, 7, 25, 23},
		{20, 11, 10, 24, 4},
		{14, 21, 16, 12, 6},
	}
	for i := range secondBoard {
		assert.ElementsMatch(secondBoard[i], bingo.Boards[1].Numbers[i], "line %v is wrong")
	}
}

func TestCheckField(t *testing.T) {
	assert := assert.New(t)

	board := NewBoard()
	board.Checked[0] = []bool{true, true, true, true, false}
	board.Numbers[0] = []int{1, 2, 3, 4, 5}
	board.Numbers[1][0] = 6
	board.SumOfUnmarked = 11
	assert.True(board.CheckField(5), "should be true")
	assert.Equal(6, board.SumOfUnmarked, "should be 6")
}

func TestSolvePart1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4512, SolvePart1("../example_data"), "wrong answer")
}
