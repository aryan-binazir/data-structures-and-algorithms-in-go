package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isSymmetric(root *TreeNode) bool {
	left, right := root.Left, root.Right
	if root == nil {
		return true
	}

	return helper(left, right)
}

func helper(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return root1 == root2
	}

	return root1.Val == root2.Val && helper(root1.Left, root2.Right) && helper(root1.Right, root2.Left)
}
