package bin

import "strconv"

func ConvBinToInt(str string) int {
	num, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
}
