package tree

// 二叉树的中序遍历，递归版本
// 中序遍历：按照访问左子树——根节点——右子树的方式遍历这棵树
func inorderTraversal(root *TreeNode) []int {
	var res []int

	// 闭包+递归
	var inorder func(r *TreeNode)
	inorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		inorder(r.Left)
		res = append(res, r.Val)
		inorder(r.Right)
	}
	return res
}

// 递归版本
func inorderTraversal2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
