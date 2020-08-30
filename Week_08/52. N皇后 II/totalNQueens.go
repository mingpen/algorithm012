package main

func main() {

}

type solution struct {
	n       int
	visited []int
	total   int
}

func (s *solution) q(visited []int, row int) {
	if row == s.n {
		s.total++
		return
	}
	for i := 0; i < s.n; i++ {
		if ok(visited, row, i) {
			visited[row] = i
			s.q(visited, row+1)
		}
	}
}

func ok(visited []int, row, col int) bool {
	rt, ct := row+col, row-col
	for i := 0; i < row; i++ {
		if visited[i] == col || i+visited[i] == rt || i-visited[i] == ct {
			return false
		}
	}
	return true
}

func totalNQueens(n int) int {
	s := &solution{n: n, visited: make([]int, n)}
	s.q(s.visited, 0)
	return s.total
}
