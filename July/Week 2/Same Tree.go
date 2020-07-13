// Link - https://leetcode.com/problems/same-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return (isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right))
	}
	return false
}

// Time - 0ms(beats 100%)
// Memory - 2.1MB(beats 100%)