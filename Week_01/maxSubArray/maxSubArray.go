func maxSubArray(nums []int) int {
	cnt := len(nums)
	if cnt == 0 {
		return 0
	}
	max, j := nums[0], 1
	for sum := max; j < cnt; j++ {
		if sum < 0 {
			sum = 0
		}
		if sum += nums[j]; sum > max {
			max = sum
		}

	}
	return max
}