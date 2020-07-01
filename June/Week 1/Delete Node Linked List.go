// Link - https://leetcode.com/problems/delete-node-in-a-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	temp := node
	parent := node
	for temp != nil && temp.Next != nil {
		temp.Val = temp.Next.Val
		parent = temp
		temp = temp.Next
	}
	parent.Next = nil
}

// Time:
//  Usage : 0ms
//  Beats : 100%
// Memory:
//  Usage : 2.9MB
//  Beats : 15.23%