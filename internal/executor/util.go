package executor

func find(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}

	return -1, false
}
