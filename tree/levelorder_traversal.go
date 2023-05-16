package tree

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	p := []*TreeNode{root}
	for i := 0; len(p) > 0; i++ {
		res = append(res, []int{})
		var q []*TreeNode
		for _, node := range p {
			if node != nil {
				res[i] = append(res[i], node.Val)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		p = q
	}
	return res
}
