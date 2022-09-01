package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return helper(root)
}

func helper(node *TreeNode) int {
	left, right := node.Left, node.Right

	if left == nil && right == nil {
		return 1
	} else if left == nil {
		return 1 + helper(right)
	} else if right == nil {
		return 1 + helper(left)
	}

	leftVal, rightVal := helper(left)+1, helper(right)+1

	if leftVal > rightVal {
		return leftVal
	}
	return rightVal
}
