package dfs

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
//此外，你可以假设该网格的四条边均被水包围。
//输入：grid = [
//["1","1","0","0","0"],
//["1","1","0","0","0"],
//["0","0","1","0","0"],
//["0","0","0","1","1"]
//]
//输出：3

// 经典的dfs|bfs
func numIslands(grid [][]byte) int {
	wide, length := len(grid), len(grid[0])
	var res int
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= wide || j < 0 || j >= length {
			// 命中为访问过的岛屿，ret +1
			return
		}
		if grid[i][j] != '1' {
			return
		}
		// 染色
		grid[i][j] = '0'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	for i := 0; i < wide; i++ {
		for j := 0; j < length; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}
