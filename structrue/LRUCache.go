package structrue

import "time"

//全文默写 & 背诵

// 普通LRU 缓存
// simple LRUCache
type LRUCache struct {
	head, tail *Node
	cap, len   int
	mem        map[int]*Node
}
type Node struct {
	next       *Node
	pre        *Node
	key, value int
}

func NewLRUCache(capacity int) LRUCache {
	l := LRUCache{}
	l.mem = map[int]*Node{}
	l.head = &Node{}
	l.tail = &Node{}
	l.cap = capacity
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) MovetoHead(node *Node) {
	if node != nil && node.pre != nil {
		node.pre.next = node.next
		node.next.pre = node.pre
	}

	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) DeleteLastOne() {
	todelete := this.tail.pre
	todelete.pre.next = this.tail
	this.tail.pre = todelete.pre
	delete(this.mem, todelete.key)
	this.len--
}
func (this *LRUCache) Get(key int) int {
	if v, ok := this.mem[key]; ok {
		this.MovetoHead(v)
		return v.value
	} else {
		return -1
	}
}
func (this *LRUCache) Put(key, val int) {
	if v, ok := this.mem[key]; ok {
		v.value = val
		this.MovetoHead(v)
		return
	} else {
		newNode := &Node{}
		newNode.key = key
		newNode.value = val
		this.mem[key] = newNode
		this.MovetoHead(newNode)
		this.len++
		if this.len > this.cap {
			this.DeleteLastOne()
		}
	}
}

//支持时间过期的 LRU 缓存
//LRUCache with deadline

type LRUCacheWithDeadLine struct {

	//....
	mem map[int]*nodeWithTTL
	TTL time.Duration
	//...
}
type nodeWithTTL struct {
	key, value int
	next, pre  *nodeWithTTL
	expire     time.Time
}

func NewLRUCacheWithDeadLine(cap int, ttl time.Duration) LRUCacheWithDeadLine {
	//...
	return LRUCacheWithDeadLine{}
}
func (this *LRUCacheWithDeadLine) MovetoHead(node *nodeWithTTL) {
	//...
	//...
	node.expire = time.Now().Add(this.TTL)
}

func (this *LRUCacheWithDeadLine) DeleteOne(node *nodeWithTTL) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCacheWithDeadLine) Get(key int) int {
	//...
	if node, ok := this.mem[key]; ok {
		if time.Now().After(node.expire) {
			this.DeleteOne(node)
			return -1
		} else {
			node.expire = time.Now().Add(this.TTL)
			this.MovetoHead(node)
			return node.value
		}
	} else {
		return -1
	}
}

func (this *LRUCacheWithDeadLine) Put(key, val int) {
	if node, ok := this.mem[key]; ok {
		if time.Now().After(node.expire) {
			node.value = val
			node.expire = time.Now().Add(this.TTL)
			this.MovetoHead(node)
		} else {
			node.value = val
			this.MovetoHead(node)
		}
	} else {
		//...
	}
}
