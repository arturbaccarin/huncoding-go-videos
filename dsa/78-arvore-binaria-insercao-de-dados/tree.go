// https://youtu.be/dbD2rmHxdUc
package main

import "fmt"

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func insert(node *TreeNode, data int) *TreeNode {
	if node == nil {
		return &TreeNode{data: data}
	}

	if data <= node.data {
		node.left = insert(node.left, data)
	} else {
		node.right = insert(node.right, data)
	}

	return node
}

func preOrder(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Printf("%d -> ", node.data)
	preOrder(node.left)
	preOrder(node.right)
}

func main() {
	var root *TreeNode

	insert(root, 12)
	insert(root, 15)
	insert(root, 11)
	insert(root, 1)
	insert(root, 9)

	preOrder(root)
}
