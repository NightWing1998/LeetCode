// Link - https://leetcode.com/problems/binary-tree-level-order-traversal-ii/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	levelTravel(root, 0, &res)
	reverseArr(&res)
	return res
}

func reverseArr(res *[][]int) {
	for i, n := 0, len(*res); i < n/2; i++ {
		(*res)[i], (*res)[n-1-i] = (*res)[n-1-i], (*res)[i]
	}
}

func levelTravel(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	if level >= len(*res) {
		*res = append(*res, make([]int, 0))
	}
	(*res)[level] = append((*res)[level], root.Val)
	levelTravel(root.Left, level+1, res)
	levelTravel(root.Right, level+1, res)
}

// Time: 0ms(beats 100%)
// Memory: 3MB(beats 51.61%)