// Link - https://leetcode.com/problems/search-in-a-binary-search-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	temp := root
	for temp != nil {
		if temp.Val == val {
			return temp
		} else if temp.Val < val {
			temp = temp.Right
		} else {
			temp = temp.Left
		}
	}
	return nil
}

// Time:
// 	Usage : 24ms
// 	Beats : 86.56%
// Memory:
// 	Usage : 6.7MB
// 	Beats : 48.72%