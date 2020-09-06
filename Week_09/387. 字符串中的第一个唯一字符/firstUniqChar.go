package main

// https://leetcode-cn.com/problems/first-unique-character-in-a-string/

func main() {

}
func firstUniqChar(s string) int {
	// 每个单词的计数
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}
	// 重新遍历字符串
	for i, c := range s {
		if m[c] == 1 {
			return i
		}
	}
	return -1
}
