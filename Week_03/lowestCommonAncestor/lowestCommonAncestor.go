package main

func main() {

}

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/submissions/

// TreeNode  .
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pA := make([]*TreeNode, 0)
	dfs(root, p, &pA)
	qA := make([]*TreeNode, 0)
	dfs(root, q, &qA)
	for _, v := range pA {
		for _, u := range qA {
			if v == u {
				return v
			}
		}
	}
	return nil
}

func dfs(root, p *TreeNode, ancestor *[]*TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p {
		*ancestor = append(*ancestor, root)
		return p
	}
	r := dfs(root.Left, p, ancestor)
	if r == p {
		*ancestor = append(*ancestor, root)
		return p
	}
	r = dfs(root.Right, p, ancestor)
	if r == p {
		*ancestor = append(*ancestor, root)
		return p
	}
	return nil
}
