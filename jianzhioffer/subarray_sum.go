package jianzhioffer

//给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。
//
//
//
//示例 1：
//
//输入:nums = [1,1,1], k = 2
//输出: 2
//解释: 此题 [1,1] 与 [1,1] 为两种不同的情况
//示例 2：
//
//输入:nums = [1,2,3], k = 3
//输出: 2
// 典型的前缀和问题
// 解释可以看：https://leetcode.cn/problems/QTMn0o/solution/qian-zhui-he-chai-zhi-de-zui-tong-su-jie-oq8q/
func subarraySum(nums []int, k int) int {
	var presum int                         // 记录当前的和
	presumCountMap := make(map[int]int, 0) // 记录当前和出现次数
	presumCountMap[0] = 1                  // 必须存储前缀和为0的情况，表示第一次presum-k=0，也就是前缀和等于k的场景
	var res int
	for _, num := range nums {
		presum += num
		res += presumCountMap[presum-k]
		presumCountMap[presum] += 1
	}
	return res
}
