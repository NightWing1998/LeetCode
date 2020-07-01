/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, 0)
}

func dfs(root *TreeNode, result int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return result*10 + root.Val
	}
	return dfs(root.Left, result*10+root.Val) + dfs(root.Right, result*10+root.Val)
}