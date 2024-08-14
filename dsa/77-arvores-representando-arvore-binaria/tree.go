package main

type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	treeNode := TreeNode{}

	treeNode.Data = "test"

	treeNode.Left = &TreeNode{
		Data: "youtube",
		Left: &TreeNode{
			Data: "go",
		},
		Right: &TreeNode{
			Data: "golang",
		},
	}
}
