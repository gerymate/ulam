package ulam

func equalMatrices(first [][]int, second [][]int) bool {
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if len(first[i]) != len(second[i]) {
			return false
		}
		for j := range first[i] {
			if first[i][j] != second[i][j] {
				return false
			}
		}
	}
	return true
}
