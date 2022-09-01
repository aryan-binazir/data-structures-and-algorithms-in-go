/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "math"

type Result struct {
	Valid  bool
	MaxLen int
}

func isBalanced(root *TreeNode) bool {
	return helper(root).Valid
}

func helper(node *TreeNode) Result {
	if node == nil {
		return Result{true, 0}
	}

	left, right := helper(node.Left), helper(node.Right)

	if !left.Valid || !right.Valid {
		return Result{false, 0}
	}

	if math.Abs(float64(left.MaxLen)-float64(right.MaxLen)) > float64(1) {
		return Result{false, 0}
	}

	if left.MaxLen > right.MaxLen {
		return Result{true, 1 + left.MaxLen}
	}
	return Result{true, 1 + right.MaxLen}
}