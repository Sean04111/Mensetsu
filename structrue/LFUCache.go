package structrue

import (
	"math"
)

//全文默写 & 背诵

type LFUCache struct {
	len, cap   int
	keymap     map[int]*LFUNode //用于查询用的
	countermap map[int]*Bucket  //这个是以counter为索引，双向链表为值
	mincounter int
}

type Bucket struct {
	head, tail *LFUNode
}
type LFUNode struct {
	key, value int
	next, pre  *LFUNode
	counter    int
}

func NewNode(key, value int) *LFUNode {
	return &LFUNode{key: key, value: value, counter: 1}
}
func NewBucket() *Bucket {
	b := &Bucket{
		head: NewNode(0, 0),
		tail: NewNode(0, 0),
	}
	b.head.next = b.tail
	b.tail.pre = b.head
	return b
}
func (this *Bucket) MovetoHead(node *LFUNode) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}
func (this *Bucket) RemoveNode(node *LFUNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}
func (this *Bucket) Isempty() bool {
	if this.head.next == this.tail {
		return true
	} else {
		return false
	}
}
func Constructor(capacity int) LFUCache {
	return LFUCache{
		len:        0,
		cap:        capacity,
		keymap:     map[int]*LFUNode{},
		countermap: map[int]*Bucket{},
		mincounter: math.MaxInt64,
	}
}
func (this *LFUCache) AddBucket(counter int) {
	this.countermap[counter] = NewBucket()
}
func (this *LFUCache) Get(key int) int {
	if node, ok := this.keymap[key]; ok {
		val := node.value
		oldbucket := this.countermap[node.counter]
		oldbucket.RemoveNode(node)
		node.counter++
		if node.counter < this.mincounter || this.countermap[this.mincounter].Isempty() {
			this.mincounter = node.counter
		} //最小counter的维护
		nextbucket, ok := this.countermap[node.counter]
		if ok {
			nextbucket.MovetoHead(node)
		} else {
			this.AddBucket(node.counter)
			this.countermap[node.counter].MovetoHead(node)
		}
		return val
	} else {
		return -1
	}
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap != 0 {
		if node, ok := this.keymap[key]; ok {
			node.value = value
			oldbucket := this.countermap[node.counter]
			oldbucket.RemoveNode(node)
			node.counter++
			if node.counter < this.mincounter || this.countermap[this.mincounter].Isempty() {
				this.mincounter = node.counter
			}
			nextbucket, ok := this.countermap[node.counter]
			if ok {
				nextbucket.MovetoHead(node)
			} else {
				this.AddBucket(node.counter)
				this.countermap[node.counter].MovetoHead(node)
			}
		} else {
			this.len++
			if this.len > this.cap {
				var opsbucket *Bucket
				if !(this.countermap[1].Isempty()) {
					opsbucket = this.countermap[1]
				} else {
					opsbucket = this.countermap[this.mincounter]
				}
				delete(this.keymap, opsbucket.tail.pre.key)
				opsbucket.RemoveNode(opsbucket.tail.pre)
				this.len--
			}
			newnode := NewNode(key, value)
			this.keymap[key] = newnode
			bucket, ok := this.countermap[1]
			if ok {
				bucket.MovetoHead(newnode)
			} else {
				this.AddBucket(1)
				this.countermap[1].MovetoHead(newnode)
			}
		}
	}
}
