package maps

func CopyMap(visits map[string]int) map[string]int {
	result := make(map[string]int, len(visits))
	for k, v := range visits {
		result[k] = v
	}
	return result
}
