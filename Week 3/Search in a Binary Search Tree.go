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