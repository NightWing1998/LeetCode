/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + height(root.Left)
}

func countNodes(root *TreeNode) int {
	rootH := height(root)
	if rootH == 0 {
		return 0
	}
	rightHeight := height(root.Right)
	if rightHeight == rootH-1 {
		// Trees have same height,i.e. last element will be in right subtree but could be in left child of right subtree hence call countNodes again on right subtree, add 2^(h-1) for left sub tree nodes as it is full
		return (1 << (rootH - 1)) + countNodes(root.Right)
	} else {
		// Height not same, means last element in left subtree, add 2^(h-2)
		return (1 << (rootH - 2)) + countNodes(root.Left)
	}
}