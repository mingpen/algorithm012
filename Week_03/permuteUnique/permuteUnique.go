package main

import "sort"

func main() {

}
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return (&solution{nums: nums, l: len(nums)}).permuteUnique()
}

type solution struct {
	// 排序
	// 子问题 nums[1:]的全排列，插入到。。。

	nums []int
	l    int
}

func (s *solution) r(nums []int) (res [][]int) {
	l := len(nums)
	if l == 0 {
		return
	}
	first := nums[0]
	if l == 1 {
		return [][]int{{first}}
	}
	arr := s.r(nums[1:])
	for _, v := range arr {
		for k := 0; k < l; k++ {
			arrNext := make([]int, l)
			res = append(res, arrNext)
			for i, j := 0, 0; i < l; i++ {
				if i == k {
					arrNext[i] = first
				} else {
					arrNext[i] = v[j]
					j++
				}
			}
			for k+1 < l && arrNext[k+1] == first {
				k++
			}
		}
	}
	return
}

func (s *solution) permuteUnique() [][]int {
	return s.r(s.nums)
}
