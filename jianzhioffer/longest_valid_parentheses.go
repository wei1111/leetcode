package jianzhioffer

//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//最大的有效括号长度，用stack
func longestValidParentheses(s string) int {
	stack := make([]int, 0, len(s))
	stack = append(stack, -1)
	var res int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				// 到边界了，重置边界
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
