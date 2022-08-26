package main

import "fmt"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	anchor, runner := head, head

	for runner != nil {
		if runner.Val != anchor.Val {
			anchor.Next = runner
			anchor = runner
		}
		runner = runner.Next
	}

	fmt.Println(anchor, runner)

	if runner != nil && anchor.Val == runner.Val {
		anchor.Next = new(ListNode)
	}
	return head
}
