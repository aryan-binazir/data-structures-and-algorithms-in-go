package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 33}
	node3 := &node{data: 16}
	node4 := &node{data: 161}
	node5 := &node{data: 126}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)

	mylist.deleteWithValue(16000) // doesn't exist
	mylist.printListData()        // expect 126 161 16 33 48
	mylist.deleteWithValue(126)   // delete start of list
	mylist.printListData()        // expect 161 16 33 48
	mylist.deleteWithValue(16)    // delete middle of list
	mylist.printListData()        // expect 161 33 48
	emptyList := linkedList{}     // empty list
	emptyList.deleteWithValue(10) // nothing should be deleted
	emptyList.printListData()     // empty list nothing should be output
}
