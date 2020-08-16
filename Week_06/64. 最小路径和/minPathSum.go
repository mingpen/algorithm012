package main

// https://leetcode-cn.com/problems/minimum-path-sum/

func main() {

}

func minPathSum(grid [][]int) int {
	rowCnt := len(grid)
	colCnt := len(grid[0])
	if rowCnt == 0 || colCnt == 0 {
		return 0
	}
	for i := 0; i < rowCnt; i++ {
		for j := 0; j < colCnt; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
				continue
			}
			if j == 0 {
				grid[i][j] += grid[i-1][j]
				continue
			}
			t := grid[i-1][j]
			if t > grid[i][j-1] {
				t = grid[i][j-1]
			}
			grid[i][j] += t
		}
	}
	return grid[rowCnt-1][colCnt-1]
}
