
package leetcode54

import (
	"testing"
)

func TestSolution(t *testing.T){
	spiralOrder([][]int{{6,9,7}})
	spiralOrder([][]int{
		{ 1, 2, 3 },
		{ 4, 5, 6 },
		{ 7, 8, 9 },
	})
}

func spiralOrder(nums [][]int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	iMin, iMax := 0, len(nums)-1
	jMin, jMax := 0, len(nums[0])-1
	ret := make([]int, 0, len(nums)*len(nums[0]))
	for iMin <= iMax && jMin <= jMax {
		// right
		for j := jMin ; j <= jMax; j ++ {
			ret = append(ret, nums[iMin][j])
		}
		// down
		for i := iMin + 1;  i <= iMax ; i ++ {
			ret = append(ret, nums[i][jMax])
		}
		// left
		if iMax > iMin {
			for j := jMax - 1; j >= jMin ; j -- {
				ret = append(ret, nums[iMax][j])
			}
		}

		// up
		if jMax > jMin{
			for i := iMax - 1; i > iMin ;  i -- {
				ret = append(ret, nums[i][jMin])
			}
		}
		jMin ++
		jMax --
		iMin ++
		iMax --

	}
	return ret
}
