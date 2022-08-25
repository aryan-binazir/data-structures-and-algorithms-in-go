package main

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	anchor, runner, previous := head, head, head

	for i := 0; i < n-1; i++ {
		runner = runner.Next
	}

	for runner.Next != nil {
		previous = anchor
		anchor, runner = anchor.Next, runner.Next
	}

	if previous == anchor {
		return anchor.Next
	}
	previous.Next = anchor.Next

	return head
}
