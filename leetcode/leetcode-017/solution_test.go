package leetcode16

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(letterCombinations("23"))
}

var m = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	return dfs(digits, "")
}

func dfs(digits string, root string) []string {
	var ret []string
	if len(digits) == 0 {
		return []string{root}
	}
	c := digits[0]
	str := m[string(c)]
	for _, s := range str {
		ret = append(ret, dfs(string(digits[1:]), root +string(s))...)
	}
	return ret
}