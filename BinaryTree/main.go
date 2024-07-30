package main

import "fmt"

type Node struct {
	data  int
	Left  *Node
	Right *Node
}

func Insert(root *Node, data int) *Node {
	if root == nil {
		return &Node{
			data:  data,
			Left:  nil,
			Right: nil,
		}
	}
	if data < root.data {
		root.Left = Insert(root.Left, data)
	} else {
		root.Right = Insert(root.Right, data)
	}
	return root
}
func Search(root *Node, data int) bool {
	if root == nil {
		return false
	}
	if root.data == data {
		return true
	}
	if data < root.data {
		return Search(root.Left, data)
	}
	return Search(root.Right, data)

}

func main() {
	root := &Node{}
	root = Insert(root, 50)
	root = Insert(root, 40)
	fmt.Println(root)
	fmt.Println(Search(root, 50))
}
