package arrays

import "AdventOfCode2021/utils/conv"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func StringArrayToIntArray(values []string) []int {
	intValues := make([]int, len(values))
	for i, value := range values {
		intValues[i] = conv.StringToNumber(value)
	}
	return intValues
}

func DeleteEmpty(values []string) []string {
	var newValues []string
	for _, value := range values {
		if value != "" {
			newValues = append(newValues, value)
		}
	}
	return newValues
}

func MakeCopy(arr []string) []string {
	newArr := make([]string, len(arr))
	copy(newArr, arr)
	return newArr
}

func PositiveNumbersSum(arr [][]int) int {
	sum := 0
	for _, row := range arr {
		for _, number := range row {
			if number > 0 {
				sum += number
			}
		}
	}
	return sum
}

func Contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func ExtractFirstValues(arr [][]int) []int {
	var values []int
	for _, row := range arr {
		if row != nil {
			values = append(values, row[0])
		}
	}
	return values
}

func MaxValue(arr []int) int {
	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

func MiddleScore(arr []int) int {
	middle := len(arr) / 2
	return arr[middle]
}

func ReverseStrings(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
