package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//思路，遍历二叉树，累加在范围内的节点val
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root==nil{
		return 0
	}
	if root.Val>R{
		return rangeSumBST(root.Left,L,R)
	}
	if root.Val<L{
		return rangeSumBST(root.Right,L,R)
	}

	return root.Val+rangeSumBST(root.Right,L,R)+rangeSumBST(root.Left,L,R)

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//思路，遍历二叉树，累加在范围内的节点val
func rangeSumBST2(root *TreeNode, L int, R int) int {
	count:=0
	if root.Right==nil&&root.Left==nil{
		if root.Val>=L&&root.Val<=R{
			return root.Val
		}
		return 0
	}
	if root.Val>=L&&root.Val<=R{
		count+=root.Val
	}
	if root.Left!=nil{
		count+=rangeSumBST(root.Left,L,R)
	}
	if root.Right!=nil{
		count+=rangeSumBST(root.Right,L,R)
	}
	return count
}