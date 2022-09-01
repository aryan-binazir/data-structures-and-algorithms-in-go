
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	result := 0

	if root.Val >= low && root.Val <= high {
		result += root.Val
	}

	if root.Val > low {
		result += rangeSumBST(root.Left, low, high)
	}

	if root.Val < high {
		result += rangeSumBST(root.Right, low, high)
	}

	return result
}