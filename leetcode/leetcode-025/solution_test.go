package leetcode25

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestSolution(t *testing.T) {
    reverseKGroup(list([]int{1,2,3,4,5}), 2)
}

func list(s []int) *ListNode {
    var ret =&ListNode{}
    end := ret
    for _, i := range s {
        end.Next = &ListNode{
            Val: i,
        }
        end = end.Next
    }
    return ret.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	ret := &ListNode{
		Next: head,
	}
	p, c := ret, head // 分别记录组的开头前一个节点和游标
	n := 0
	for c != nil {
		n++
		if n%k != 0 {
			c = c.Next
			continue
		} else {
            // 需要保留的指针  c.Next
            // 这么写太复杂了， 可以把reverse的逻辑单独抽出来
            // cNext := c.Next
            // newHead := c.Next
            // oldHead := p.Next
            // p.Next = c
            // p = oldHead
            // for oldHead != cNext {
            //     next := oldHead.Next
            //     oldHead.Next = newHead
            //     newHead = oldHead
            //     oldHead = next
            // }
            // c = cNext
            

            // 说明
            // 保留三个指针，c.Next p.Next和p
            cNext := c.Next
            c.Next = nil
            pNext := p.Next
            // 然后翻转
            rList := reverseList(pNext)
            // 将翻转之后的连进链表
            p.Next = rList
            pNext.Next = cNext
            // 更新父节点和游标位置
            p = pNext
            c = cNext
		}
	}
	return ret.Next
}

// 翻转链表
func reverseList(head *ListNode) *ListNode {
    var ret *ListNode = nil
    for head != nil {
        next := head.Next
        head.Next = ret
        ret = head
        head = next
    }
    return ret
}
