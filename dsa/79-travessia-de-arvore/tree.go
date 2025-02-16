// https://youtu.be/7szQY0iAQD0
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

// Printa o dado atual e depois vai para os filhos
func preOrder(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Printf("%d -> ", node.data)
	preOrder(node.left)
	preOrder(node.right)
}

// todos os nós de esquerda e depois da direita
func inOrder(node *TreeNode) {
	if node == nil {
		return
	}

	inOrder(node.left)
	fmt.Printf("%d -> ", node.data)
	inOrder(node.right)
}

func postOrder(node *TreeNode) {
	if node == nil {
		return
	}

	postOrder(node.left)
	postOrder(node.right)
	fmt.Printf("%d -> ", node.data)
}

func main() {
	var root *TreeNode

	root = insert(root, 10)
	insert(root, 2)
	insert(root, 0)
	insert(root, 20)
	insert(root, 3)
	insert(root, 4)

	preOrder(root)
}
