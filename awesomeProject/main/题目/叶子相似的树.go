package 题目

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 只要两颗树的叶子节点的排序相同就判定为相似的树
// 递归
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	l1 := sortTree(root1)
	l2 := sortTree(root2)
	if len(l1) != len(l2) {
		return false
	}
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}
func sortTree(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		result = append(result, root.Val)
		return result
	} else {
		result = append(result, sortTree(root.Left)...)

		result = append(result, sortTree(root.Right)...)
	}
	return result
}
