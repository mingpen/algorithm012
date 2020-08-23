package main

import (
	"log"
	"sort"
)

func main() {
	// println(jump([]int{2, 3, 1, 1, 4}))
	log.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(nums []int, target int) (res [][]int) {
	sort.Ints(nums)
	// 暴力法
	resMap := make(map[[4]int]bool)
	lenNums := len(nums)
	for i := 0; i < lenNums-3; i++ {
		iV := nums[i]
		for j := i + 1; j < lenNums-2; j++ {
			jV := nums[j]
			for k := j + 1; k < lenNums-1; k++ {
				kV := nums[k]
				for l := k + 1; l < lenNums; l++ {
					lV := nums[l]
					if iV+jV+kV+lV == target {
						resMap[[4]int{iV, jV, kV, lV}] = true
					}
				}
			}
		}
	}

	for k := range resMap {
		v := make([]int, 4)
		copy(v, k[:])
		res = append(res, v)
	}
	return
}

func jump(nums []int) int {
	index1 := len(nums) - 1
	cnts := make([]int, len(nums))
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		return 1
	}
	lenOfNums := len(nums)
	cnts[index1] = 0
	index1--
	cnts[index1] = 1
	for index1--; index1 >= 0; index1-- {
		maxStep := nums[index1]
		if maxStep >= lenOfNums-index1-1 {
			cnts[index1] = 1
			continue
		}
		//
		var min = lenOfNums - index1 - 1
		maxIndex := len(nums)
		if maxIndex > maxStep+index1+1 {
			maxIndex = maxStep + index1 + 1
		}
		for i := index1 + 1; i < maxIndex; i++ {
			if cnts[i] < min {
				min = cnts[i]
			}
			// 小于等于1步没有不要再循环
			if min == 1 {
				break
			}
		}
		cnts[index1] = min + 1
	}
	return cnts[0]
}

func jump2(nums []int) int {
	// 边界条件
	if len(nums) <= 1 {
		return 0
	}
	if len(nums) == 2 {
		return 1
	}
	maxStep := nums[0]
	if len(nums) <= maxStep+1 {
		return 1
	}
	// 穷举法
	var min = len(nums) - 1
	for i := 1; i-1 < maxStep; i++ {
		t := jump2(nums[i:])
		if t < min {
			min = t
		}
	}
	return min + 1
}
