package main

import (
	"fmt"

	. "./list"
)

func goDeeper(node *TreeNode, preorder *[]int) {
	if node != nil {
		goDeeper(node.Left, preorder)
		goDeeper(node.Right, preorder)
		*preorder = append(*preorder, node.Val)
	}
}

func postorderTraversal(root *TreeNode) []int {
	res := &[]int{}
	goDeeper(root, res)
	return *res
}

func main() {
	fmt.Println(
		postorderTraversal(NewTree([]interface{}{1, "null", 2, 3})),
	)
}
