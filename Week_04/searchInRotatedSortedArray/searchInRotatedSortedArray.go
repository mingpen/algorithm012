package main

func main() {

}
func search(nums []int, target int) int {
	start, end := 0, len(nums)
	for start < end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[start] <= nums[end-1] && target > nums[mid] ||
			nums[mid] >= nums[start] && (target > nums[mid] || target < nums[mid] && target < nums[start]) ||
			nums[mid] < nums[start] && target > nums[mid] && target < nums[start] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return -1
}
