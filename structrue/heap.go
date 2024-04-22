package structrue

// 小根堆实现
type Heap []int

func (h Heap) Up(i int) {
	for {
		j := (i - 1) / 2
		if j == i || h[j] < h[i] {
			break
		}
		h[i], h[j] = h[j], h[i]
		i = j
	}
}

func (h Heap) Down(i int) {
	for {
		leftChild := i*2 + 1
		if leftChild >= len(h) {
			break
		}
		next := leftChild
		if rightChild := i*2 + 2; rightChild < len(h) && h[rightChild] < h[next] {
			next = rightChild
		}
		if h[next] > h[i] {
			break
		}
		h[next], h[i] = h[i], h[next]
		i = next
	}
}

func (h *Heap) Pop() int {
	ans := (*h)[0]
	(*h)[0], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[0]
	*h = (*h)[:len(*h)-1]
	h.Down(0)
	return ans
}

func (h *Heap) Push(x int) {
	*h = append(*h, x)
	h.Up(len(*h) - 1)
}
