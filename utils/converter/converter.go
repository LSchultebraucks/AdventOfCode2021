package converter

import "strconv"

func StringToNumber(value string) int {
	intVar, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return intVar
}
