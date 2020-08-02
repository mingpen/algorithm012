package main

func main() {
	search([]int{4, 5, 6, 7, 0, 1, 2})
}

func search(nums []int) int {
	start, end := 0, len(nums)
	if end == 0 || nums[end-1] >= nums[start] {
		return -1
	}
	for start < end {
		mid := (start + end) / 2
		if nums[mid] >= nums[0] { // 上段
			start = mid + 1
		} else {
			end = mid
		}
	}
	if nums[start] > nums[0] {
		start++
	}
	println(nums[start], start)
	return start
}
