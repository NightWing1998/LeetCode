// Link - https://leetcode.com/problems/remove-linked-list-elements/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	for ; head != nil && head.Val == val; head = head.Next {
	}
	if head == nil {
		return nil
	}
	for parent, temp := head, head.Next; temp != nil; temp = temp.Next {
		if temp.Val == val {
			parent.Next = temp.Next
		} else {
			parent = temp
		}
	}
	return head
}

// Time - 8ms(beats 87.30%)
// Memory - 4.7MB(beats 100%)