package main

// https://leetcode-cn.com/problems/walking-robot-simulation/description/

func main() {

}

func robotSim(commands []int, obstacles [][]int) int {

	return 0
}

type solution struct {
	commands  []int //
	obstacles [][]int
	//
	direction string // +y [初始状态], -y, +x , -x

	point struct {
		X int
		Y int
	} // 最新点的位置

	maxDistance int
}
