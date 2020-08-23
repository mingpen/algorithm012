package main

func main() {

}

// https://leetcode-cn.com/problems/n-queens/
type solution struct {
	res [][]string
	n   int
}

func (s *solution) solveNQueens(v []int, row int) {
	if row == s.n {
		// 已经填充完了
		s.print(v)
		return
	}
	for i := 0; i < s.n; i++ {
		v[row] = i // row 行，i列
		if !ok(v, row, i) {
			continue
		}
		s.solveNQueens(v, row+1)
	}
}

func (s *solution) print(v []int) {
	arr := make([]string, 0, s.n)
	for _, col := range v {
		str := ""
		for i := 0; i < s.n; i++ {
			if i == col {
				str += "Q"
			} else {
				str += "."
			}
		}
		arr = append(arr, str)
	}
	s.res = append(s.res, arr)
}

func ok(v []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if v[i] == col || v[i]+i == col+row || v[i]-i == col-row {
			return false
		}
	}
	return true
}

func solveNQueens(n int) [][]string {
	s := &solution{n: n}
	s.solveNQueens(make([]int, n), 0)
	return s.res
}

//
