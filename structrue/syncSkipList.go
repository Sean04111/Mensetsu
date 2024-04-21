package structrue

import "sync"

//其他保持一致
//node 中加一把锁
type SyncNode struct{
	nexts []*SyncNode
	locker sync.RWMutex
	key,val int
}

//在 get、put 操作的时候获取前一个节点的节点锁即可
