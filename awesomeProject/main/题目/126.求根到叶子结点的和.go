package 题目

//type TreeNode struct {
//	     Val int
//	     Left *TreeNode
//	     Right *TreeNode
//	 }
// 从根开始每一个节点带一个数组
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left, right int
	result := root.Val
	if root.Right == nil && root.Left == nil {
		return result
	} else {
		if root.Left != nil {
			root.Left.Val += root.Val * 10
			left = sumNumbers(root.Left)
		}
		if root.Right != nil {
			root.Right.Val += root.Val * 10
			right = sumNumbers(root.Right)
		}

	}
	return left + right
}
