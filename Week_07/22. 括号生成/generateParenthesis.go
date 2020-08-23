package main

func main() {

}

func generateParenthesis(n int) []string {
	s := &solution{n: n}
	s.gen("", 0, 0)
	return s.res
}

type solution struct {
	res []string
	n   int
}

func (s *solution) gen(str string, left, right int) {
	if right == s.n {
		s.res = append(s.res, str)
		return
	}
	if left < s.n {
		s.gen(str+"(", left+1, right)
	}
	if left > right {
		s.gen(str+")", left, right+1)
	}
}
