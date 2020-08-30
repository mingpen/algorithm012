package main

// https://leetcode-cn.com/problems/number-of-1-bits/

func main() {

}

func hammingWeight(num uint32) (cnt int) {
	for ; num > 0; cnt++ {
		num &= num - 1
	}
	return
}
