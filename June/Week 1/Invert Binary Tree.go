// Link - https://leetcode.com/problems/invert-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		root.Val,
		invertTree(root.Right),
		invertTree(root.Left),
	}
}

// Time:
//  Usage : 0ms
//  Beats : 100%
// Memory:
//  Usage : 2.2MB
//  Beats : 12.29%