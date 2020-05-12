package alibaba1

import (
	"container/heap"
	"fmt"
	"testing"
)

// 最大堆
type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] } // 最大堆用>, 最小堆用<
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestHeap(t *testing.T) {
	h := &Heap{2, 1, 5, 100, 3, 6, 1, 0}
	heap.Init(h)
	// for _, item := range []int{2, 1, 5, 100, 3, 6, 4, 5} {
	// heap.Push(h, item)
	// }
	for h.Len() > 0 {
		val := heap.Pop(h)
		t.Logf("%d", val)
	}
}

func TestChicken(t *testing.T) {
	var m, n, k int
	store := make(Heap, n)

	fmt.Scanf("%d %d %d", &n, &m, &k)
	for i := 0; i < n; i++ {
		fmt.Scan("%d", store[i])
	}

	n, m, k = 3, 3, 100
	store = []int{100, 200, 400}
	solution(n, m, k, store)
}

// O(n * m)
func rawSolution(n, m, k int, store []int) {
	for d := 1; d <= m; d++ {
		// 增加操作
		// 注意增加并不会改变相对关系
		// 也就是说增加前最大的，增加后仍然是最大的
		maxIdx := 0

		// 这里2n复杂度可以合成1n
		for i := 0; i < len(store); i++ {
			if store[i] > store[maxIdx] {
				maxIdx = i
			}
		}

		for i := 0; i < len(store); i++ {
			store[i] += k
		}

		// 除二操作
		store[maxIdx] /= 2

		// 所以实质上， 只需要找到最大值， 计算出模拟的减少值
		// store[maxIdx] = (store[maxIdx] + i * k )
		//  =>
		// store[maxIdx] = (store[maxIdx]+k) / 2
		//  =>
		// store[maxIdx] = store[maxIdx] - (store[maxIdx]- i*k) / 2
	}

	sum := 0
	for _, i := range store {
		sum += i
	}
	fmt.Printf("%d\n", sum)
}

func solution(n, m, k int, store Heap) {
	h := &store
	heap.Init(h)

	// 每日迭代
	// 从1开始
	// 通过递推，每一次减少的量是  (tmp - i * d) / 2  主要是负数的情况如何考虑呢
	for d := 1; d <= m; d++ {
		// 获取堆顶， 也就是最大数
		tmp := heap.Pop(h).(int)

		// 最大数减去 d * k
		// 为什么这里是减去
		// 因为相比于正常情况下， 这里的tmp已经加过d轮，原来应该是 tmp + d * k， 除以2后减去的值是（tmp + d * k ）/ 2
		// 那么这里的减量也需要模拟出这种情况，这里实际上 tmp = tmp -  （tmp + d*k） / 2 => tmp = (tmp - d*k) / 2
		tmp -= d * k

		// 处理负数情况？
		if tmp%2 != 0 {
			tmp--
		}

		// 取走一半
		tmp /= 2

		heap.Push(h, tmp)
	}

	sum := 0
	for h.Len() > 0 {
		sum += heap.Pop(h).(int)
	}
	sum += m * k * n
	fmt.Printf("%d\n", sum)
}
