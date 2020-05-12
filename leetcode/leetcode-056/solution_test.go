package leetcode056

import (
	"testing"
	"sort"
)
func TestSolution(t *testing.T) {

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {	
	// 合并区间的关键在于需要有上下界重合的地方
	// 可以尝试根据下界对数组排序
	// 然后一次遍历合并
	var ret [][]int
	if len(intervals) == 0 {
		return ret
	}
	sort.Slice(intervals, func(i, j int)bool {
		return intervals[i][0] < intervals[j][0]
	})

	cur := intervals[0]
	for i := 1 ; i < len(intervals); i ++{
		if intervals[i][0] <= cur[1] { // 合并
			cur[1] = max(intervals[i][1], cur[1])
		} else {
			ret = append(ret, cur)
			cur = intervals[i]
		}
	}
	ret = append(ret, cur)
	return  ret
}