package leetcode057

import (
	"testing"
)
func TestSolution(t *testing.T){
	insert([][]int{{1,5}}, []int{2, 7})
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func insert(intervals [][]int, newInterval []int) [][]int {
	var ret [][]int
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	var cur  []int
	for i := 0 ; i < len(intervals) ; i ++ {
		cur = intervals[i]
		if newInterval == nil {
			ret = append(ret, cur)
		} else if cur[1] < newInterval[0]{ // 无重合
			ret = append(ret, cur)
		} else if cur[0] > newInterval[1] {
			ret = append(ret, newInterval)
			newInterval = nil
			ret = append(ret, cur)
		}else { // 有重合, 那么进行区间合并
			newInterval[0] = min(newInterval[0], cur[0])
			newInterval[1] = max(newInterval[1], cur[1])
		}
	}
	if newInterval != nil {
		ret = append(ret, newInterval)
	}
	return ret
}