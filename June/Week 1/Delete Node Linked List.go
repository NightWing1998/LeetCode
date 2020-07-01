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