package main

//import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var prev *TreeNode

func validate(root *TreeNode, low *int, high *int) bool {
	//Empty Tree is a Valid tree
	if root == nil {
		return true
	}

	//Check current node
	if (low != nil && root.Val <= *low) || (high != nil && root.Val >= *high) {
		return false
	}
	//return validate Left & Right node
	return validate(root.Right, &root.Val, high) && validate(root.Left, low, &root.Val)

}

func isValidBST(root *TreeNode) bool {
	res := validate(root, nil, nil)
	return res
}

func inorder(root *TreeNode) bool {

	if root == nil {
		return true
	}

	if !inorder(root.Left) {
		return false
	}

	if prev != nil && root.Val <= prev.Val {
		return false
	}
	prev = root

	return inorder(root.Right)
}

func isValidBSTinorder(root *TreeNode) bool {
	prev = nil
	return inorder(root)
}

func main() {

}
