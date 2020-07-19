package main

func main() {
	b := isInterleave(
		"aabcc",
		"dbbca",
		"aadbbcbcac")
	println(b)
}
func isInterleave(s1 string, s2 string, s3 string) bool {
	//
	s1l, s2l, s3l := len(s1), len(s2), len(s3)
	p1, p2 := 0, 0
	for i := 0; i < s3l; {
		return p1 < s1l && s1[p1] == s3[i] && isInterleave(s1[p1+1:], s2, s3[1:]) ||
			p2 < s2l && s2[p2] == s3[i] && isInterleave(s1, s2[1:], s3[1:])
	}
	return s1 == "" && s2 == ""
}
