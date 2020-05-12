
package leetcode29

import (
	"testing"
    "fmt"
    "math"
)

func TestSolution(t *testing.T){
	// fmt.Println(divide(1, 2))
	// fmt.Println(divide(7, 3))
    // fmt.Println(divide(7, -3))
    fmt.Println(divide(10, 10))
}

func divide(dividend int, divisor int) int {
    var val []int


    sign := (dividend >= 0 && divisor > 0)  || (dividend <0 && divisor < 0)
    
    if dividend == int(math.MinInt32) && (divisor == -1) {
        return int(math.MaxInt32)
    }


    dividend = abs(dividend)
    divisor = abs(divisor)

    sum := divisor
    for dividend >= sum {
        val = append(val, sum)
        sum += sum
    }

	var ret int
    for index := len(val) -1 ; index >= 0 ; index -- {
        ret = ret << 1
		if dividend >= val[index] {
            dividend -= val[index]
			ret ++
		}
		
    }
    if sign {
        return ret
    } 
    return 0-ret
}

func abs(a int) int {
    if a < 0 {
        return 0-a
    }
    return a
}