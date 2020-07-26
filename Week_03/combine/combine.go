package main

func main() {

}

func combine(n int, k int) [][]int {
	return (&solution{n: n, k: k}).combine()
}

type solution struct {
	res [][]int
	k   int
	n   int
}

func (s *solution) r(start, level int, arr []int) {
	if level == s.k {
		s.res = append(s.res, arr)
		return
	}
	for i := start + 1; i <= s.n; i++ {
		arrDrill := make([]int, s.k)
		copy(arrDrill, arr)
		arrDrill[level] = i
		s.r(i, level+1, arrDrill)
	}
}
func (s *solution) combine() [][]int {
	//
	level := 0
	for i := 1; i <= s.n-s.k+1; i++ {
		arr := make([]int, s.k)
		arr[0] = i
		s.r(i, level+1, arr)
	}
	return s.res
}
