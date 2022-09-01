package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}

	currentNode := root
	// stack = append(stack, root)

	for currentNode != nil {
		stack = append(stack, currentNode)          // push currentNode to stack
		lastValInStack := stack[len(stack)-1]       // pop; get last val
		stack = stack[:len(stack)-1]                // pop; reset stack to new length
		result = append(result, lastValInStack.Val) // pop from stack, and push to result
		if currentNode.Right != nil {
			stack = append(stack, currentNode.Right)
		}
		currentNode = currentNode.Left
	}

	for len(stack) > 0 {
		lastValInStack := stack[len(stack)-1] // pop; get last val
		stack = stack[:len(stack)-1]          // pop; reset stack to new length
		currentNode := lastValInStack
		result = append(result, currentNode.Val) // pop from stack, and push to result
		if currentNode.Right != nil {
			stack = append(stack, currentNode.Right)
		}
		currentNode = currentNode.Left
	}
	return result
}
