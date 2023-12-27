package main


func findMirror(grid []string) int {
	for r := 1; r < len(grid); r++ {
		above := reverseStrSlice(grid[:r])
		below := grid[r:]

		// Ensure that the length of above is not greater than the length of below
		if len(above) > len(below) {
			above = above[:len(below)]
		}

		below = below[:len(above)]

		if equalSlices(above, below) {
			return r
		}
	}

	return 0
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func rotate90(matrix []string) []string {
	result := make([]string, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		var row string
		for j := len(matrix) - 1; j >= 0; j-- {
			row += string(matrix[j][i])
		}
		result[i] = row
	}
	return result
}


