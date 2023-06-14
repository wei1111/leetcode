package dfs

import (
	"testing"
)

func Test_permute(t *testing.T) {
	nums := []int{1, 2, 3}
	got := permute(nums)
	t.Logf("%v", got)
}
