package myleetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	l1 := new(ListNode)
	l2 := new(ListNode)
	z1 := l1
	z2 := l2
	t := head
	i := 1
	for t != nil {
		if i%2 == 1 {
			z1.Next = t
			z1 = z1.Next
		} else {
			z2.Next = t
			z2 = z2.Next
		}
		t = t.Next
		i++
	}
	if i%2 == 0 {
		z2.Next = nil
	} else {
		z1.Next = nil
	}
	z1.Next = l2.Next
	return l1.Next
}
