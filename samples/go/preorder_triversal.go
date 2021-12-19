package main

import (
	"fmt"

	. "./list"
)

func goDeeper(node *TreeNode, preorder *[]int) {
	if node != nil {
		*preorder = append(*preorder, node.Val)
		goDeeper(node.Left, preorder)
		goDeeper(node.Right, preorder)
	}
}

func preorderTraversal(root *TreeNode) []int {
	res := &[]int{}
	goDeeper(root, res)
	return *res
}

func main() {
	fmt.Println(
		preorderTraversal(NewTree([]interface{}{1, "null", 2, 3})),
	)
}
