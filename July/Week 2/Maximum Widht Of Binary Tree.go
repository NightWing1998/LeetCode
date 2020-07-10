// Link - https://leetcode.com/problems/maximum-width-of-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Depth struct {
	root  *TreeNode
	depth int
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]Depth, 0, 1)
	result := 1
	queue = append(queue, Depth{root: root, depth: 0})
	for len(queue) > 0 {
		s := queue[0].depth
		e := queue[len(queue)-1].depth
		result = max(result, e-s+1)
		newQueue := make([]Depth, 0, len(queue)*2)
		for _, node := range queue {
			index := node.depth - s
			if node.root.Left != nil {
				newQueue = append(newQueue, Depth{root: node.root.Left, depth: 2*index + 1})
			}
			if node.root.Right != nil {
				newQueue = append(newQueue, Depth{root: node.root.Right, depth: 2*index + 2})
			}
			queue = newQueue
		}
	}
	return result
}

// Time - 0ms(beats 100%)
// Memory - 3.3MB(beats 91.67%)