package tree

func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return res
}

func preorderTraversal2(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	for len(stack) > 0 || root != nil {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		// POP
		root = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		root = root.Right
	}
	return res
}
