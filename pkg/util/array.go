package util

func IntInArray(search int, arr []int) bool {
	for _, val := range arr {
		if val == search {
			return true
		}
	}

	return false
}
