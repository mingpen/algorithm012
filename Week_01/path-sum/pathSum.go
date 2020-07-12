package main

import "log"

func main() {

	log.Println(hasPathSum(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
	}, 22))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	// 非递归前序遍历
	var parent *TreeNode
	var top *StackNode
	for root != nil {
		// 先对自己赋值
		if parent != nil {
			root.Val += parent.Val
		}
		leaf := root.Left == nil && root.Right == nil
		if leaf && root.Val == sum {
			return true
		}
		if root.Right != nil {
			// 右节点入栈
			root.Right.Val += root.Val
			newTop := new(StackNode)
			newTop.Val = root.Right
			newTop.next, top = top, newTop
		}
		if root.Left != nil {
			// 左叶子非空
			root, parent = root.Left, root
			continue
		}
		// 出栈
		if top == nil {
			break
		}
		parent = nil
		root, top = top.Val, top.next
	}
	return false
}

// 右节点,才会入栈,入栈的值包含了路径值
type StackNode struct {
	Val  *TreeNode
	next *StackNode
}
