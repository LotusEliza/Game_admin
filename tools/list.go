package tools

func InListInt(i int, lst []int) bool {
	for _, it := range lst {
		if it == i {
			return true
		}
	}
	return false
}
