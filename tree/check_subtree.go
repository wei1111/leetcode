package tree

//判断二叉树是否是子树
//检查子树。你有两棵非常大的二叉树：T1，有几万个节点；T2，有几万个节点。设计一个算法，判断 T2 是否为 T1 的子树。
//如果 T1 有这么一个节点 n，其子树与 T2 一模一样，则 T2 为 T1 的子树，也就是说，从节点 n 处把树砍断，得到的树与 T2 完全相同。
func checkSubTree(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val == right.Val && isEqual(left, right) {
		return true
	} else {
		return checkSubTree(left.Left, right) || checkSubTree(left.Right, right)
	}
}

func isEqual(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isEqual(left.Left, right.Left) && isEqual(left.Right, right.Right)
}

//通过二叉树的遍历，可以将树转换成字符串，其中前序遍历如果将空节点表示出来，那么树就是唯一的。
//时间复杂度：O(n+m)，n和m分别是t1和t2中节点的数目。
//空间复杂度：O(n+m)
