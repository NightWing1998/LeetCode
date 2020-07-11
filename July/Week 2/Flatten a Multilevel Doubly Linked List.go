// Link - https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	flatChild(root)
	return root
}

// This function returns the last child in the list
func flatChild(root *Node) *Node {
	temp := root
	parent := root
	for temp != nil {
		parent = temp
		next := temp.Next
		if temp.Child != nil {
			temp.Next = temp.Child
			temp.Child.Prev = temp
			lastNode := flatChild(temp.Child)
			temp.Child = nil
			if next != nil {
				lastNode.Next = next
				next.Prev = lastNode
			}
		}
		temp = next
	}
	return parent
}

// Time - 0ms(beats 100%)
// Memory - 2.9MB(beats 80%)