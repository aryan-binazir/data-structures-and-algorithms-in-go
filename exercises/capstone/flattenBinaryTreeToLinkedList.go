package main

func flatten(root *TreeNode) {
	arr := preorderTraversal(root)

	for i := 0; i < len(arr); i++ {
		node := arr[i]
		node.Left = nil
		if i == len(arr)-1 {
			node.Right = nil
		} else {
			node.Right = arr[i+1]
		}
	}
}

func preorderTraversal(root *TreeNode) []*TreeNode {
	result := []*TreeNode{}
	stack := []*TreeNode{}

	currentNode := root
	// stack = append(stack, root)

	for currentNode != nil {
		stack = append(stack, currentNode)      // push currentNode to stack
		lastValInStack := stack[len(stack)-1]   // pop; get last val
		stack = stack[:len(stack)-1]            // pop; reset stack to new length
		result = append(result, lastValInStack) // pop from stack, and push to result
		if currentNode.Right != nil {
			stack = append(stack, currentNode.Right)
		}
		currentNode = currentNode.Left
	}

	for len(stack) > 0 {
		lastValInStack := stack[len(stack)-1] // pop; get last val
		stack = stack[:len(stack)-1]          // pop; reset stack to new length
		currentNode := lastValInStack
		result = append(result, currentNode) // pop from stack, and push to result
		if currentNode.Right != nil {
			stack = append(stack, currentNode.Right)
		}
		if currentNode.Left != nil {
			stack = append(stack, currentNode.Left)
		}
		// currentNode = currentNode.Left
	}
	return result
}
