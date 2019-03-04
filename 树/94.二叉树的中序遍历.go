package main

// 二叉树转换成中序链表，不能用额外空间，可以递归
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
//示例:
//
//输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [1,3,2]
// 先输出左边的再输出中间的，再输出右边的

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}

	if root.Left != nil {
		lcount := inorderTraversal(root.Left)
		result = append(result, lcount...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		rcount := inorderTraversal(root.Right)
		result = append(result, rcount...)
	}
	return result
}
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{}
	result := []int{}
	for len(queue) != 0 || root != nil {
		for root != nil {
			queue = append(queue, root)
			root = root.Left
		}
		if len(queue) != 0 {
			node := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			result = append(result, node.Val)
			root = node.Right

		}

	}
	return result
}

// List<Integer> list = new ArrayList<Integer>();
//
//        Stack<TreeNode> s = new Stack<TreeNode>();
//
//        do {
//            while (root != null) {
//                s.push(root);
//                root = root.left;
//            }
//            if (!s.isEmpty()) {
//                TreeNode node = s.pop();
//                list.add(node.val);
//                root = node.right;
//            }
//        } while (!s.isEmpty() || root != null);
