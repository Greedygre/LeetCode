package main

//给定一个二叉树，原地将它展开为链表。
//
//例如，给定二叉树
//例如，给定二叉树
//
//    1
//   / \
//  2   5
// / \   \
//3   4   6
//将其展开为：
//
//1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6
// 记得把左边置零
// 思路1：先跟读取之后全部压入一个队列。不符合原地修改。
// 思路2，先序遍历，root的左子树所有节点都在右子树前边。
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	if root.Left != nil {

		right := root.Right
		root.Right = root.Left
		root.Left = nil
		node := root.Right
		for node.Right != nil {
			node = node.Right
		}
		node.Right = right
	}

}
