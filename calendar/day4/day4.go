package main

import (
	"AdventOfCode2021/utils/arrays"
	"AdventOfCode2021/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(4)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0
	numbersDrawn := arrays.StringArrayToIntArray(strings.Split(input[0], ","))
	bingoBoards := make([][][]int, (len(input)-1)/6)
	bingoBoardIdx := 0
	for idx := 2; idx < len(input); idx += 6 {
		bingoBoards[bingoBoardIdx] = readBingoBoard(input, idx)
		bingoBoardIdx++
	}
	idxBoard, number := drawNumbers(numbersDrawn, bingoBoards)
	winnerBingoBoardSum := arrays.PositiveNumbersSum(bingoBoards[idxBoard])
	result = number * winnerBingoBoardSum
	return result
}

func solvePart2(input []string) int {
	result := 0
	numbersDrawn := arrays.StringArrayToIntArray(strings.Split(input[0], ","))
	bingoBoards := make([][][]int, (len(input)-1)/6)
	bingoBoardIdx := 0
	for idx := 2; idx < len(input); idx += 6 {
		bingoBoards[bingoBoardIdx] = readBingoBoard(input, idx)
		bingoBoardIdx++
	}
	result = drawNumbersLastToWin(numbersDrawn, bingoBoards)
	return result
}

func drawNumbersLastToWin(numbersDrawn []int, bingoBoards [][][]int) int {
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
					return boardsWithBingo[bingoBoardsWithBingoCnt-1][1]
				}
			}
		}
	}
	return boardsWithBingo[bingoBoardsWithBingoCnt-1][1]
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

func drawNumbers(numbersDrawn []int, bingoBoards [][][]int) (int, int) {
	for _, number := range numbersDrawn {
		bingoBoards = markNumberInBingoBoards(number, bingoBoards)
		idxBoard := checkForBingo(bingoBoards)
		if idxBoard != -1 {
			return idxBoard, number
		}
	}
	panic("No bingo found")
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

func checkForBingo(bingoBoards [][][]int) int {
	for idx := range bingoBoards {
		if checkForBingoInRows(bingoBoards[idx]) || checkForBingoInColumns(bingoBoards[idx]) {
			return idx
		}
	}
	return -1
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
