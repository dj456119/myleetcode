/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-14 18:15:09
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-14 19:20:58
 */
/*
 * @Descripttion:给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。

 * @version:
 * @Author: cm.d
 * @Date: 2021-10-14 18:15:09
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-14 18:15:43
*/

package myleetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Left == nil {
		return root
	}
	root.Left.Next = root.Right
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}

	connect(root.Left)
	connect(root.Right)
	return root
}
func Connect(root *Node) *Node {
	return connect(root)
}
