package structrue

import "testing"

func TestHeap(t *testing.T){
	h:=Heap{}
	for i:=100;i>=0;i--{
		h.Push(i)
	}
	t.Log(h.Pop())
}
