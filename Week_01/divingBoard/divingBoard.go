package main

import "log"

func main() {
	log.Println(divingBoard(1, 1, 0))
}
func divingBoard(shorter int, longer int, k int) []int {
	// k = 0
	if k == 0 {
		return []int{}
	}
	// shorter 与 longer相等
	if shorter == longer {
		return []int{shorter * k}
	}

	res := make([]int, k+1)
	// i 从0 到k,表示多少个longer
	for i := 0; i < k+1; i++ {
		res[i] = (k-i)*shorter + i*longer
	}
	return res
}
