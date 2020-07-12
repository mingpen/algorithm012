package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 保存进数
	var more int
	var head, pre *ListNode
	for l1 != nil || l2 != nil || more > 0 {
		v := more
		more = 0
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		if v >= 10 {
			more, v = 1, v-10
		}
		p := new(ListNode)
		p.Val = v
		if pre != nil {
			pre.Next, pre = p, p
		}
		if head == nil {
			pre, head = p, p
		}
	}
	//
	return head
}
