package main

// https://leetcode-cn.com/problems/lemonade-change/
func main() {

}

func lemonadeChange(bills []int) bool {
	// 记录 5 10 的个数
	m := map[int]int{}
	for _, v := range bills {
		switch v {
		case 5:
			m[5]++
		case 10:
			m[10]++
			c := m[5]
			if c == 0 {
				return false
			}
			m[5]--
		case 20:
			c5 := m[5]
			c10 := m[10]
			if c10 > 0 {
				if c5 == 0 {
					return false
				}
				m[10]--
				m[5]--
			} else {
				if c5 < 3 {
					return false
				}
				m[5] -= 3
			}
		}
	}
	return true
}
