package leetcodem1708

import (
	"testing"
	"sort"
)


func TestSolution(t *testing.T){
	// bestSeqAtIndex([]int{65,70,56,75,60,68}, []int{100,150,90,190,95,110})
	bestSeqAtIndex([]int{2868,5485,1356,1306,6017,8941,7535,4941,6331,6181},
		 []int{5042,3995,7985,1651,5991,7036,9391,428,7561,8594})
}

type size [][2]int
func (s size) Less(i, j int) bool { return (s[i][0] < s[j][0]) || (s[i][0] == s[j][0] && s[i][1] > s[j][1]) }
func (s size) Swap(i, j int) { s[i], s[j] = s[j], s[i]}
func (s size) Len() int {return len(s)}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int{
	if a < b {
		return b
	}
	return a
}

func bestSeqAtIndex(height []int, weight []int) int {
	l := min(len(height), len(weight))
	s := make(size, l)
	for i := 0 ; i < l ; i ++ {
		s[i][0], s[i][1] = height[i], weight[i]
	} 
	sort.Sort(s)
	
	// 按高度求最大上升子序列
	dp := make([]int, l)
	for i := 0 ; i < l ; i ++ {
		m := 1
		for j := 0 ; j < i ; j ++ {
			hi, hj, wi, wj :=  s[i][0], s[j][0],s[i][1],s[j][1]
			if hi > hj && wi > wj {
				m = max(m, dp[j] + 1)  
			}
		}
		dp[i] = m
	}
	
	return dp[l-1]
	return ret
}

func bestSeqAtIndex2(height []int, weight []int) int {
	length := min(len(height), len(weight))
	s := make([][2]int, length)
	for i := 0 ; i < length ; i ++ {
		s[i][0], s[i][1] = height[i], weight[i]
	} 
	sort.Slice(s, func(i, j int) bool {return s[i][0] < s[j][0] || (s[i][0] == s[j][0] && s[i][1] > s[j][1]) })
	
	// 按高度求最大上升子序列
	dp := make([]int, 1, length) // dp里面存的是最小体重
	dp[0] = s[0][1]
	for i := 1 ; i < length ; i ++ {
		l, r := 0, len(dp) - 1
		if dp[r] < s[i][1] {
			dp = append(dp, s[i][1])
			continue
		}
		for l < r {
			m := (l + r) / 2
			if dp[m] < s[i][1] {
				l = m + 1
			} else {
				r = m
			}
		}
		dp[l] = s[i][1]
	}
	
	return len(dp)
}