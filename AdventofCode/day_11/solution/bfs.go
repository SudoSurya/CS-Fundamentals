package main

import (
	"fmt"
)

func  shortestPath(grid [][]string, source, destination [2]int) int {
	// Edge Case: if the source is only the destination.
	if source[0] == destination[0] && source[1] == destination[1] {
		return 0
	}

	// Create a queue for storing cells with their distances from source
	// in the form {dist,{cell coordinates pair}}.
	queue := [][]int{{0, source[0], source[1]}}
	n, m := len(grid), len(grid[0])

    fmt.Println(n, m)
    fmt.Println("source", source)
    fmt.Println("destination", destination)
	// Create a distance matrix with initially all the cells marked as
	// unvisited and the source cell as 0.
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, m)
		for j := range dist[i] {
			dist[i][j] = 1e9
		}
	}
	dist[source[0]][source[1]] = 0

	// The following delta rows and delta columns array are created such that
	// each index represents each adjacent node that a cell may have
	// in a direction.
	dr := []int{-1, 0, 1, 0}
	dc := []int{0, 1, 0, -1}

	// Iterate through the maze by popping the elements out of the queue
	// and pushing whenever a shorter distance to a cell is found.
	for len(queue) > 0 {
		it := queue[0]
		queue = queue[1:]
		dis, r, c := it[0], it[1], it[2]

		// Through this loop, we check the 4 direction adjacent nodes
		// for a shorter path to destination.
		for i := 0; i < 4; i++ {
			newr, newc := r+dr[i], c+dc[i]

			// Checking the validity of the cell and updating if dist is shorter.
			if newr >= 0 && newr < n && newc >= 0 && newc < m && grid[newr][newc] == "1" && dis+1 < dist[newr][newc] {
				dist[newr][newc] = 1 + dis

				// Return the distance until the point when
				// we encounter the destination cell.
				if newr == destination[0] && newc == destination[1] {
					return dis + 1
				}
				queue = append(queue, []int{1 + dis, newr, newc})
			}
		}
	}
	// If no path is found from source to destination.
	return -1
}

func BFS() {
	// Driver Code.
	source := [2]int{0, 1}
	destination := [2]int{2, 2}

	grid := [][]string{
		{"1", "1", "1", "1"},
		{"1", "1", "0", "1"},
		{"1", "1", "1", "1"},
		{"1", "1", "0", "0"},
		{"1", "0", "0", "1"},
	}


	res := shortestPath(grid, source, destination)

	fmt.Println(res)
}
