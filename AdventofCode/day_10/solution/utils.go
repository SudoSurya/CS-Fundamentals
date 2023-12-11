package main

func PipeNextPos(pipes [][]string, pos []int, prevPos []int, direction string) []int {
	if direction == "up" {
		return getUpPos(pipes, pos, prevPos)
	}
	return getDownPos(pipes, pos, prevPos)
}

func dashNextPos(pipes [][]string, pos []int, prevPos []int, direction string) []int {
	if direction == "left" {
		return getLeftPos(pipes, pos, prevPos)
	}
	return getRightPos(pipes, pos, prevPos)
}

func LNextPos(pipes [][]string, pos []int, prevPos []int, direction string) []int {
	if direction == "up" {
		return getUpPos(pipes, pos, prevPos)
	}
	return getRightPos(pipes, pos, prevPos)
}

func JnextPos(pipes [][]string, pos []int, prevPos []int, direction string) []int {
	if direction == "up" {
		return getUpPos(pipes, pos, prevPos)
	}
	return getLeftPos(pipes, pos, prevPos)
}

func SevenNextPos(pipes [][]string, pos []int, prevPos []int, direction string) []int {
	if direction == "left" {
		return getLeftPos(pipes, pos, prevPos)
	}
	return getDownPos(pipes, pos, prevPos)
}

func FNextPos(pipes [][]string, pos []int, prevPos []int, direction string) []int {
	if direction == "right" {
        return getRightPos(pipes, pos, prevPos)
	}
	return getDownPos(pipes, pos, prevPos)
}

func getUpPos(pipes [][]string, pos []int, prevPos []int) []int {
	if pos[0] > 0 {
		if notDot(pipes, []int{pos[0] - 1, pos[1]}) && notSamePos([]int{pos[0] - 1, pos[1]}, prevPos) {
			return []int{pos[0] - 1, pos[1]}
		}
	}
	return []int{-1, -1}
}

func getDownPos(pipes [][]string, pos []int, prevPos []int) []int {
	if pos[0] < len(pipes)-1 {
		if notDot(pipes, []int{pos[0] + 1, pos[1]}) && notSamePos([]int{pos[0] + 1, pos[1]}, prevPos) {
			return []int{pos[0] + 1, pos[1]}
		}
	}
	return []int{-1, -1}
}

func getLeftPos(pipes [][]string, pos []int, prevPos []int) []int {
	if pos[1] > 0 {
		if notDot(pipes, []int{pos[0], pos[1] - 1}) && notSamePos([]int{pos[0], pos[1] - 1}, prevPos) {
			return []int{pos[0], pos[1] - 1}
		}
	}
	return []int{-1, -1}
}

func getRightPos(pipes [][]string, pos []int, prevPos []int) []int {
	if pos[1] < len(pipes[0])-1 {
		if notDot(pipes, []int{pos[0], pos[1] + 1}) && notSamePos([]int{pos[0], pos[1] + 1}, prevPos) {
			return []int{pos[0], pos[1] + 1}
		}
	}
	return []int{-1, -1}
}

func notSamePos(pos []int, prevPos []int) bool {
	return pos[0] != prevPos[0] || pos[1] != prevPos[1]
}
