package structrue

import (
	"math/rand"
)

//impliment of skiplist(no concurrency)

type SkipList struct {
	head *node
}

// use int type key & val
type node struct {
	nexts    []*node
	key, val int
}

func NewSkipList()*SkipList{
	return &SkipList{
		head:&node{},
	}
}

func (s *SkipList) search(target int) *node {
	curr := s.head
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		for curr.nexts[level] != nil && curr.nexts[level].key < target {
			curr = curr.nexts[level]
		}
		if curr.nexts[level] != nil && curr.nexts[level].key == target {
			return curr.nexts[level]
		}
	}
	return nil
}

func (s *SkipList) roll() int {
	level := 0
	for rand.Intn(2) > 0 {
		level++
	}
	return level
}

func (s *SkipList) Get(key int) (int, bool) {
	val := s.search(key)
	if val == nil {
		return -1, false
	} else {
		return val.val, true
	}
}

func (s *SkipList) Put(key, val int) {
	if _node := s.search(key); _node != nil {
		_node.val = val
		return
	}
	curr := s.head

	//随机出新节点的层数
	maxLevel := s.roll()

	//补充 head 节点的高度
	for len(s.head.nexts)-1 < maxLevel {
		s.head.nexts = append(s.head.nexts, nil)
	}

	newNode := &node{
		key:   key,
		val:   val,
		nexts: make([]*node, maxLevel+1),
	}

	//插入新节点
	for level := maxLevel; level >= 0; level-- {
		for curr.nexts[level] != nil && curr.nexts[level].key < key {
			curr = curr.nexts[level]
		}
		newNode.nexts[level] = curr.nexts[level]
		curr.nexts[level] = newNode
	}
}

func (s *SkipList) Del(key int) {
	if _node := s.search(key); _node == nil {
		return
	}

	curr := s.head
	for level := len(curr.nexts) - 1; level >= 0; level-- {
		for curr.nexts[level] != nil && curr.nexts[level].key < key {
			curr = curr.nexts[level]
		}

		if curr.nexts[level] != nil && curr.nexts[level].key == key {
			curr.nexts[level] = curr.nexts[level].nexts[level]
		}

	}
}
