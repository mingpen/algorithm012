package main

func main() {

}

// https://leetcode-cn.com/problems/insufficient-nodes-in-root-to-leaf-paths/
/*
 * Definition for a binary tree node.
 */

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return root
	}
	dfs(&root, limit, 0)
	return root
}

func dfs(proot **TreeNode, limit, sum int) {
	root := *proot
	sum += root.Val
	leaf := root.Left == nil && root.Right == nil
	if leaf && sum < limit {
		*proot = nil
	}
	if leaf {
		return
	}

	// 非叶子节点
	if root.Left != nil {
		dfs(&root.Left, limit, sum)
	}
	if root.Right != nil {
		dfs(&root.Right, limit, sum)
	}
	if root.Left == nil && root.Right == nil {
		*proot = nil
	}
}
