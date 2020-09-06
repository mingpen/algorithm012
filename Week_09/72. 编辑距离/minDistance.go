package main

// https://leetcode-cn.com/problems/edit-distance/

func main() {
	minDistance("dinitrophenylhydrazine",
		"benzalphenylhydrazone")
}

func minDistance(word1 string, word2 string) int {
	// return (&solution{
	// 	m:  make(map[[2]int]int),
	// 	s1: word1, s2: word2}).dp(len(word1), len(word2))
	return (&solution{
		m:  make(map[[2]int]int),
		s1: word1, s2: word2}).dp2()
}

//
type solution struct {
	s1, s2 string
	m      map[[2]int]int
	status [][]int
}

func (s *solution) dp2() int {
	l1 := len(s.s1)
	l2 := len(s.s2)
	// 构建数组
	s.status = make([][]int, l1+1)
	count := l2 + 1
	for i := range s.status {
		s.status[i] = make([]int, count)
		for j := 0; j < count; j++ {
			if j == 0 {
				s.status[i][0] = i
			} else if i == 0 {
				s.status[0][j] = j
				// 下面的步骤，i 和 j 均大于0
			} else if s.s1[i-1] == s.s2[j-1] {
				s.status[i][j] = s.status[i-1][j-1]
			} else {
				s.status[i][j] = 1 + min(s.status[i-1][j-1], s.status[i][j-1], s.status[i-1][j])
			}
		}
	}
	return s.status[l1][l2]
}

func (s *solution) dp(i, j int) (v int) {
	defer func() {
		s.m[[2]int{i, j}] = v
	}()
	if v, ok := s.m[[2]int{i, j}]; ok {
		return v
	}
	if i == 0 {
		return j
	}
	if j == 0 {
		return i
	}
	// i ,j 均大于0
	if s.s1[i-1] == s.s2[j-1] {
		return s.dp(i-1, j-1)
	}
	return 1 + min(s.dp(i-1, j-1), s.dp(i, j-1), s.dp(i-1, j))
}

func min(x, y, z int) int {
	if x < y {
		y = x
	}
	if y < z {
		z = y
	}
	return z
}
