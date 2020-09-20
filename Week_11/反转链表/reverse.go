package main

import "log"

// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/

/**
 * Definition for singly-linked list.
 */

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode // 始终指向 head 的前一个值
	for head.Next != nil {
		pre, head, head.Next = head, head.Next, pre
	}
	// 到达终止条件, head.Next 为nil 了，需要指向 pre
	head.Next = pre
	return head
}

func reverseList2(cur *ListNode) *ListNode {
	var pre *ListNode // 始终指向 cur 的前一个值
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	// cur 为nil,循环终止
	return pre
}

func reverseList4(head *ListNode) *ListNode {
	h, _ := reverseList3(head)
	return h
}

func reverseList3(head *ListNode) (head2, tail2 *ListNode) {
	if head.Next == nil {
		return head, head
	}
	h2, t2 := reverseList3(head.Next)
	head2 = h2
	tail2, t2.Next, head.Next = head, head, nil
	return
}

func reverseListK(head *ListNode, k int) *ListNode {

	var cur *ListNode = head
	var tmpHead *ListNode = cur

	i := 0
	// 移动1个就包含2个node，移动2个包含3个node
	for ; cur != nil && i < k-1; i++ {
		cur = cur.Next
	}
	// 2. 截取链表；反转
	if i < k-1 || cur == nil {
		// 停止
		return head
	}
	nextHead := cur.Next
	cur.Next = nil // 截断
	head = reverseList(tmpHead)

	// 3. 链接: 反转后 头变尾;
	tmpHead.Next = reverseListK(nextHead, k)

	return head
}

// https://github.com/lifei6671/interview-go/blob/master/question/q016.md#%E9%A2%98%E7%9B%AE
// 变形的链表反转题

func reverseListK2(head *ListNode, k int) *ListNode {
	if k < 2 {
		// k 小于2，无需操作
		return head
	}
	// 1. 反转
	head = reverseListK(reverseList(head), k)
	// 2. 再次反转
	return reverseList(head)
}

//

func main() {
	var head = &ListNode{Val: 1}
	var cur *ListNode = head
	for i := 2; i < 9; i++ {
		cur.Next = &ListNode{
			Val: i,
		}
		cur = cur.Next
	}
	printList(reverseListK2(head, 3))
}

func printList(head *ListNode) {
	var arr []int
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	log.Println(arr)
}
