package leetcode039

import (
	"bytes"
	"fmt"
	"testing"
	"sort"
)

func TestSolution(t *testing.T){
	// combinationSum([]int{2,3,6,7}, 7)
	// combinationSum([]int{2,3,5}, 8)
	combinationSum([]int{2,3,7}, 18)
}


type Set map[string][]int

func (s Set)Add(args []int){
	var dst = make([]int, len(args))
	copy(dst, args)
	buf := bytes.NewBufferString("")
	for _, i := range args {
		buf.WriteString(fmt.Sprint(i))
		buf.WriteString("+")
	}
	s[buf.String()] = dst
}

// time O(len(candidates)*target)
func combinationSum(candidates []int, target int) [][]int {
	var dp = make(map[int]Set)
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	for i := candidates[0]; i <= target ; i ++ {
		st := make(Set)
		for _, j := range candidates {
			if i == j {
				st.Add([]int{i})
			} else if i > j {
				if _, exist := dp[i-j]; exist {
					for _, s := range dp[i-j]{
						newSeq := append(s,j)
						sort.Ints(newSeq)
						st.Add(newSeq)
					}
				}
			}
		}
		if len(st) > 0 {
			dp[i] = st
		}
	}
	var ret = make([][]int, 0, len(dp[target]))
	for _, val := range dp[target] {
		ret = append(ret, val)
	}
	return ret
}