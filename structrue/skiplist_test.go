package structrue

import "testing"

func TestBasic(t *testing.T) {
	s:=NewSkipList()
	s.Put(1,1)
	s.Del(1)
	t.Log(s.Get(1))
}