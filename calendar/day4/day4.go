package main

import (
	"AdventOfCode2021/utils/arrays"
	"AdventOfCode2021/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(4)
	println(SolvePart1(input))
	println(SolvePart2(input))
}

func SolvePart1(input []string) int {
	result := 0
	numbersDrawn, bingoBoards := parseInput(input)
	boardsWithBingo := playBingo(numbersDrawn, bingoBoards)
	result = boardsWithBingo[0][1]
	return result
}

func SolvePart2(input []string) int {
	result := 0
	numbersDrawn, bingoBoards := parseInput(input)
	boardsWithBingo := playBingo(numbersDrawn, bingoBoards)
	result = boardsWithBingo[len(boardsWithBingo)-1][1]
	return result
}

func parseInput(input []string) ([]int, [][][]int) {
	numbersDrawn := arrays.StringArrayToIntArray(strings.Split(input[0], ","))
	bingoBoards := make([][][]int, (len(input)-1)/6)
	bingoBoardIdx := 0
	for idx := 2; idx < len(input); idx += 6 {
		bingoBoards[bingoBoardIdx] = readBingoBoard(input, idx)
		bingoBoardIdx++
	}
	return numbersDrawn, bingoBoards
}

func playBingo(numbersDrawn []int, bingoBoards [][][]int) [][]int {
	boardsWithBingo := make([][]int, len(bingoBoards))
	totalBingoBoards := len(bingoBoards)
	bingoBoardsWithBingoCnt := 0
	for _, number := range numbersDrawn {
		bingoBoards = markNumberInBingoBoards(number, bingoBoards)
		for i := 0; i < totalBingoBoards; i++ {
			excludeIdx := arrays.ExtractFirstValues(boardsWithBingo)
			idxBoard := checkForBingoExclude(bingoBoards, excludeIdx)
			if idxBoard != -1 && notAlreadyAddedAsBingo(boardsWithBingo, idxBoard) {
				boardsWithBingo[bingoBoardsWithBingoCnt] = make([]int, 2)
				boardsWithBingo[bingoBoardsWithBingoCnt][0] = idxBoard
				boardsWithBingo[bingoBoardsWithBingoCnt][1] = number * arrays.PositiveNumbersSum(bingoBoards[idxBoard])
				bingoBoardsWithBingoCnt++
				if bingoBoardsWithBingoCnt == totalBingoBoards {
					return boardsWithBingo
				}
			}
		}
	}
	return boardsWithBingo[:totalBingoBoards]
}

func notAlreadyAddedAsBingo(boardsWithBingo [][]int, idx int) bool {
	isAdded := false
	for _, s := range boardsWithBingo {
		if s != nil && s[0] == idx {
			isAdded = true
		}
	}
	return !isAdded
}

func markNumberInBingoBoards(number int, bingoBoards [][][]int) [][][]int {
	for i, bingoBoard := range bingoBoards {
		for j, row := range bingoBoard {
			for k, numberDrawn := range row {
				if numberDrawn == number {
					bingoBoards[i][j][k] = -1
				}
			}
		}
	}
	return bingoBoards
}

func checkForBingoExclude(bingoBoards [][][]int, exclude []int) int {
	for idx := range bingoBoards {
		if !arrays.Contains(exclude, idx) &&
			(checkForBingoInRows(bingoBoards[idx]) || checkForBingoInColumns(bingoBoards[idx])) {
			return idx
		}
	}
	return -1
}

func checkForBingoInColumns(board [][]int) bool {
	for i := 0; i < len(board); i++ {
		isBingo := true
		for j := 0; j < len(board); j++ {
			if board[j][i] != -1 {
				isBingo = false
				break
			}
		}
		if isBingo {
			return true
		}
	}
	return false
}

func checkForBingoInRows(board [][]int) bool {
	for i := 0; i < len(board[0]); i++ {
		isBingo := true
		for j := 0; j < len(board); j++ {
			if board[i][j] != -1 {
				isBingo = false
				break
			}
		}
		if isBingo {
			return true
		}
	}
	return false
}

func readBingoBoard(input []string, idx int) [][]int {
	bingoBoard := make([][]int, 5)
	for i := 0; i < 5; i++ {
		bingoBoard[i] = readBingoLine(input[idx+i])
	}
	return bingoBoard
}

func readBingoLine(line string) []int {
	return arrays.StringArrayToIntArray(arrays.DeleteEmpty(strings.Split(line, " ")))
}
