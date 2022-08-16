package main

import "fmt"

// Node represents the components of a binary tree
type Node struct {
	key   int
	left  *Node
	right *Node
}

// Insert will add node to the tree
func (n *Node) Insert(key int) {
	if n.key < key {
		// move right
		if n.right == nil {
			n.right = &Node{key: key}
		} else {
			n.right.Insert(key)
		}
	} else if n.key > key {
		// move left
		if n.left == nil {
			n.left = &Node{key: key}
		} else {
			n.left.Insert(key)
		}
	}
}

// Search will return true if there is a node with that value
func (n *Node) Search(key int) bool {

	if n == nil {
		return false
	}
	if n.key < key {
		// move right
		return n.right.Search(key)
	} else if n.key > key {
		// move left
		return n.left.Search(key)
	}
	// if not smaller than or larger than, match
	return true
}

func main() {
	tree := &Node{key: 100}
	// fmt.Println(tree)
	tree.Insert(300)
	tree.Insert(200)
	tree.Insert(301)
	tree.Insert(20)
	tree.Insert(33)
	tree.Insert(1)
	tree.Insert(1000)
	tree.Insert(950)
	tree.Insert(75)
	tree.Insert(513)
	tree.Insert(99)
	tree.Insert(833)
	tree.Insert(834)
	tree.Insert(123)

	fmt.Println(tree.Search(200))  // true
	fmt.Println(tree.Search(834))  // true
	fmt.Println(tree.Search(123))  // true
	fmt.Println(tree.Search(1))    // true
	fmt.Println(tree.Search(2))    // false
	fmt.Println(tree.Search(2000)) // false
	fmt.Println(tree.Search(831))  // false
	fmt.Println(tree.Search(222))  // false
}
