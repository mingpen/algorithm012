package main

import "log"

func main() {
	nums := []int{1, 0, 11, 23}
	moveZeroes(nums)
	log.Println(nums)
}

// 双指针
func moveZeroes(nums []int) {
	// 边界条件
	cnt := len(nums)
	if cnt < 1 {
		return
	}
	var nonZero, i int = 0, 0
	for ; i < cnt; i++ {
		if nums[i] != 0 {
			if nonZero == i { // 同一个数无需复制
				nonZero++
				continue
			}
			nums[nonZero], nums[i] = nums[i], 0
			nonZero++
		}
	}
}
