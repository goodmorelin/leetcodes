/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * tosee: https://leetcode.cn/problems/insertion-sort-list/submissions/
 */
func insertionSortList(head *ListNode) *ListNode {
  help := &ListNode{}
	cur := head
	tail := help.Next
	for cur != nil {
		if tail != nil && cur.Val > tail.Val {
			tail.Next = cur
			cur = cur.Next
			tail = tail.Next
            tail.Next = nil
			continue
		}
		tmp := help
		for tmp.Next != nil && cur.Val > tmp.Next.Val {
			tmp = tmp.Next
		}
		tmp2 := cur
		cur = cur.Next
		tmp2.Next = tmp.Next
		tmp.Next = tmp2
		if tail == nil {
			tail = tmp.Next
		}
	}
	return help.Next
}
