package dfs

// 全排列
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	temp := make([]int, len(nums))
	onPath := make([]bool, len(nums))
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			res = append(res, append([]int{}, temp...))
		}
		for j, on := range onPath {
			if !on {
				temp[i] = nums[j]
				onPath[j] = true
				dfs(i + 1)
				onPath[j] = false
			}

		}
	}

	dfs(0)
	return res
}
