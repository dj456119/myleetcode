/*
 * @Descripttion:请编写一个函数，用于 删除单链表中某个特定节点 。在设计函数时需要注意，你无法访问链表的头节点 head ，只能直接访问 要被删除的节点 。

题目数据保证需要删除的节点 不是末尾节点

 * @version:
 * @Author: cm.d
 * @Date: 2021-11-02 21:00:58
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-02 21:05:12
*/

package myleetcode

type ListNode237 struct {
	Val  int
	Next *ListNode237
}

func deleteNode(node *ListNode237) {
	next := node.Next
	node.Val, next.Val = next.Val, node.Val
	node.Next = next.Next
}

func DeleteNode(node *ListNode237) {
	deleteNode(node)
}
