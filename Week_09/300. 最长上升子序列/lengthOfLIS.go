package main

// https://leetcode-cn.com/problems/longest-increasing-subsequence/
func main() {

}

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	dp := make([]int, length)
	dp[0] = 1
	for i := 1; i < length; i++ {
		dp[i] = 1
		for j, v := range nums[:i] {
			if v < nums[i] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	max := dp[0]
	for _, v := range dp[1:] {
		if v > max {
			max = v
		}
	}
	return max
}
