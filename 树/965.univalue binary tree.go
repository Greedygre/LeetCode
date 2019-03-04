package main

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	uval := root.Val
	if root.Right != nil && root.Right.Val != uval {
		return false
	}

	if root.Left != nil && root.Left.Val != uval {
		return false
	}
	rcount := isUnivalTree(root.Right)
	lcount := isUnivalTree(root.Left)
	if rcount && lcount {
		return true
	}
	return false
}

//class Solution {
//    public boolean isUnivalTree(TreeNode root) {
//    if (root == null) {
//            return true;
//        }
//        Set<Integer> res = new HashSet<>();
//        Queue<TreeNode> queue = new LinkedList<>();
//        queue.add(root);
//        while (!queue.isEmpty()) {
//            int len = queue.size();
//            for (int i = 0; i < len; i++) {
//                TreeNode node = queue.poll();
//                res.add(node.val);
//                if (node.left != null) {
//                    queue.add(node.left);
//                }
//                if (node.right != null) {
//                    queue.add(node.right);
//                }
//            }
//        }
//        return res.size() == 1 ? true : false;
//    }
//
//}
func isUnivalTree2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	uval := root.Val
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		len := len(queue)
		for i := 0; i < len; i++ {
			node := queue[0]
			queue = queue[1:]
			if uval != node.Val {
				return false
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return true
}
