package main

func main() {

}

// https://leetcode-cn.com/problems/permutations/submissions/

func permute(nums []int) (res [][]int) {
	//
	l := len(nums)
	if l == 0 {
		return nil
	}
	first := nums[0]
	if l == 1 {
		return [][]int{{first}}
	}
	before := permute(nums[1:])
	for _, arr := range before {
		l := len(arr)
		for i := 0; i < l+1; i++ {
			arrNext := make([]int, l+1)
			arrNext[i] = first
			res = append(res, arrNext)
			for j, k := 0, 0; k < l; j++ {
				if i == j {
					continue
				}
				arrNext[j] = arr[k]
				k++
			}
		}
	}
	return

}
