package main

func main() {

}

// https://leetcode-cn.com/problems/friend-circles/
func findCircleNum(M [][]int) int {
	s := &solution{M: M}
	return s.findCircle()
}

type solution struct {
	M       [][]int
	n       int
	visited []bool
}

func (s *solution) findCircle() (cnt int) {
	s.n = len(s.M)
	s.visited = make([]bool, s.n)
	for i := 0; i < s.n; i++ {
		if !s.visited[i] {
			cnt++
			s.dfs(i)
		}
	}
	return
}

func (s *solution) dfs(v int) {
	if s.visited[v] {
		return
	}
	s.visited[v] = true
	for i := 0; i < s.n; i++ {
		if i != v && s.M[v][i] == 1 {
			s.dfs(i)
		}
	}
}
