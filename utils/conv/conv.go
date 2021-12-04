package conv

import (
	"strconv"
	"strings"
)

func StringToNumber(value string) int {
	value = strings.Trim(value, " ")
	intVar, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return intVar
}
