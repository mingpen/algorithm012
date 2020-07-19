package main

import "log"

func main() {
	r := minimumTotal(arr)
	log.Println(r)
}

func minimumTotal(triangle [][]int) int {
	height := len(triangle)
	for h := height - 1; h > 0; h-- {
		for i := 0; i < h; i++ {
			min := triangle[h][i]
			if min > triangle[h][i+1] {
				min = triangle[h][i+1]
			}
			triangle[h-1][i] += min
		}
	}
	return triangle[0][0]
}

func minimumTotal2(triangle [][]int) int {
	// 二维数组对应一个满二叉树
	// i 行(非最后一行),j列的 孩子为 [i+1,j],[i+1,j+1]
	// 转化为二叉树最小路径和
	// 前序遍历 （根，左，右）
	// pathsum
	p := &pathSum{notSet: true, maxLin: len(triangle) - 1}
	p.pathsum(triangle, [2]int{0, 0}, 0)

	return p.min
}

type pathSum struct {
	min    int
	notSet bool
	maxLin int
}

func (p *pathSum) pathsum(triangle [][]int, index [2]int, before int) {
	sum := before + triangle[index[0]][index[1]]
	if index[0] == p.maxLin {
		// 叶子节点
		if p.notSet || sum < p.min {
			p.notSet = false
			p.min = sum
		}
		return
	}
	p.pathsum(triangle, [2]int{index[0] + 1, index[1]}, sum)
	p.pathsum(triangle, [2]int{index[0] + 1, index[1] + 1}, sum)
}
