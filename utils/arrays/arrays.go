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

func MakeCopy(arr []string) []string {
	newArr := make([]string, len(arr))
	copy(newArr, arr)
	return newArr
}
