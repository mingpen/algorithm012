package main

// https://leetcode-cn.com/problems/sliding-window-maximum/

import (
	"container/list"
	"log"
)

func main() {

	log.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
func maxSlidingWindow(nums []int, k int) []int {

	// 单调递减队列，队头是最大的元素 【入队时可确定最大的值】
	// 最大值出队时，重新寻找最大值

	l := list.New()
	var res []int

	cnt := len(nums)
	if k > cnt {
		k = cnt
	}
	// 初始化
	for i := 0; i < k; i++ {
		push(l, nums, i)
	}
	if e := l.Front(); e != nil {
		res = append(res, nums[e.Value.(int)])
	}
	//
	for i := k; i < cnt; i++ {
		push(l, nums, i)

		fe := l.Front()
		f := fe.Value.(int)
		if f == i-k { // 淘汰数据
			l.Remove(fe)
			f = l.Front().Value.(int)
		}
		res = append(res, nums[f])

	}

	return res
}

func push(l *list.List, nums []int, i int) {
	for e := l.Back(); e != nil && nums[i] > nums[e.Value.(int)]; {
		e, _ = e.Prev(), l.Remove(e)
	}
	l.PushBack(i)
}
