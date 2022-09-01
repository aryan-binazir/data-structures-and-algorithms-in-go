package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return helper(p, q)
}

func helper(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	return p.Val == q.Val && helper(p.Left, q.Left) && helper(p.Right, q.Right)
}
