package main

import "fmt"

type TreeNode struct {
	val       int
	leftNode  *TreeNode
	RightNode *TreeNode
}

func findHeight(root *TreeNode) int {

	// base condition > smallest valid input is null node ht
	if root == nil {
		return 0
	}

	// find ht of left or right sub trees recursively
	leftTree := findHeight(root.leftNode)
	rightTree := findHeight(root.RightNode)

	if leftTree > rightTree {
		return leftTree + 1
	}

	return rightTree + 1
}

func main() {

	// Creating a sample binary tree:
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5

	root := &TreeNode{val: 1}
	root.leftNode = &TreeNode{val: 2}
	root.RightNode = &TreeNode{val: 3}
	root.leftNode.leftNode = &TreeNode{val: 4}
	root.leftNode.RightNode = &TreeNode{val: 5}

	fmt.Println(findHeight(root))

}
