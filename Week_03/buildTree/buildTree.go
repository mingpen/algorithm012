package main

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 105. 从前序与中序遍历序列构造二叉树

func main() {

}

/**
 * Definition for a binary tree node.

 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	lPre, lIn := len(preorder), len(inorder)
	if lPre == 0 || lPre != lIn {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if lPre == 1 {
		return root
	}
	rootIndex := 0
	for ; rootIndex < lIn; rootIndex++ {
		if root.Val == inorder[rootIndex] {
			break
		}
	}
	if rootIndex == lIn {
		return nil //异常情况
	}
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
