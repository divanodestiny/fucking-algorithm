package leetcode040

import (
	"bytes"
	"fmt"
	"testing"
	"sort"
)

func TestSolution(t *testing.T){
	// combinationSum([]int{2,3,6,7}, 7)
	// combinationSum([]int{2,3,5}, 8)
	combinationSum2([]int{10,1,2,7,6,1,5}, 8)
}


type Set map[string][]int

func (s Set)Add(args []int){
	var dst = make([]int,0, len(args))
	buf := bytes.NewBufferString("")
	for _, i := range args {
		if i > 0 {
			buf.WriteString(fmt.Sprint(i))
			buf.WriteString("+")
			dst = append(dst, i)
		}
	}
	s[buf.String()] = dst
}

// time O(len(candidates)*target)
func combinationSum2(candidates []int, target int) [][]int {
	var set = make(Set)
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)

	var backtrace func(s []int, i int)

	backtrace= func(s []int, i int) {
		if i == len(s){
			return 
		}
		total := sum(s[0:i+1])
		if s[i] == 2 || s[i] == 6 {
			s[i] = s[i]
		}
		if total == target {
			set.Add(s[0:i+1])
			return 
		} else if total > target {
			return 
		} else {
			val := s[i]
			s[i] = 0
			backtrace(s, i+1)
			s[i] = val
			backtrace(s, i+1)
		}
		
	}
	backtrace(candidates, 0)

	var ret = make([][]int, 0, len(set))
	for _, val := range set {
		ret = append(ret, val)
	}
	return ret
}

func sum(args []int) int {
	var ret int 
	for _, i := range args {
		 ret += i
	}
	return ret 
}

