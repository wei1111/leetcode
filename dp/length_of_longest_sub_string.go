package dp

//状态定义: 设动态规划列表 dp ，dp[j] 代表以字符 s[j] 为结尾的“最长不重复子字符串”的长度
//转移方程: 固定右边界j，设字符 s[j] 左边距离最近的相同字符为s[i]，即s[i] = s[j].
//1.当i< 0，即 s[j] 左边无相同字符，则 dp[j] = dp[j-1] +1
//2.当 dp[j -1]<j-i，说明字符 s[i] 在子字符串 dp[j-1] 区间之外则 dp[j] = dp[j-1] +1
//3.当 dp[j -1]>=j-i，说明字符 s[i] 在子字符串 dp[j -1] 区间之中则 dp[j]; 的左边界由 s[i] 决定，即 dp[j] = j-i
//当i< 0时，由于 dp[j -1]<=j恒成立，因而 dp[j -1] < j-i恒成立因此分支 1. 和 2. 可被合并。
// dp[j] = dp[j-1]+1, dp[j-1]<j-i
// dp[j] = j-i, dp[j-1]>=j-i

// 最长不重复子字符串
func lengthOfLongestSubstring(s string) int {
	var res, tmp int
	charMap := make(map[byte]int, len(s))
	for j, _ := range s {
		i := -1
		if _, ok := charMap[s[j]]; ok {
			i = charMap[s[j]]
		}
		charMap[s[j]] = j

		if tmp < j-i {
			tmp += 1
		} else {
			tmp = j - i
		}

		res = max(tmp, res)
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
