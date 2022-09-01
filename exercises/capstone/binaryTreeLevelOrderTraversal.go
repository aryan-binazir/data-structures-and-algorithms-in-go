package main

func levelOrder(root *TreeNode) (result [][]int) {
	cache := map[int][]int{}
	currentLevel := 0
	helper(root, currentLevel, cache)
	i := 0
	for true {
		if arr, ok := cache[i]; ok {
			result = append(result, arr)
			i++
		} else {
			break
		}
	}
	return result
}

func helper(currentNode *TreeNode, currentLevel int, cache map[int][]int) {
	if currentNode == nil {
		return
	}

	if _, ok := cache[currentLevel]; ok {
		cache[currentLevel] = append(cache[currentLevel], currentNode.Val)
	} else {
		cache[currentLevel] = []int{currentNode.Val}
	}
	helper(currentNode.Left, currentLevel+1, cache)
	helper(currentNode.Right, currentLevel+1, cache)
}
