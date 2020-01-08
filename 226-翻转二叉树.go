package main

/*
翻转一棵二叉树。
输入：
     4
   /   \
  2     7
 / \   / \
1   3 6   9

输出：
     4
   /   \
  7     2
 / \   / \
9   6 3   1
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right = invertTree(root.Right)
	root.Left = invertTree(root.Left)
	root.Left, root.Right = root.Right, root.Left
	return root
}
