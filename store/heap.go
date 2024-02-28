package store

import (
    "container/heap"
    "time"
)

type ItemHeap []*Item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].Expiration < h[j].Expiration }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i]; h[i].Index = i; h[j].Index = j }

func (h *ItemHeap) Push(x interface{}) {
    n := len(*h)
    item := x.(*Item)
    item.Index = n
    *h = append(*h, item)
}

func (h *ItemHeap) Pop() interface{} {
    old := *h
    n := len(old)
    item := old[n-1]
    item.Index = -1 
    *h = old[0 : n-1]
    return item
}

func (h *ItemHeap) Update(item *Item, value interface{}, expiration int64) {
    item.Value = value
    item.Expiration = expiration
    heap.Fix(h, item.Index)
}

func InitHeap() *ItemHeap {
    var h ItemHeap
    heap.Init(&h)
    return &h
}

func (h *ItemHeap) CleanupExpiredItems() {
    now := time.Now().UnixNano()
    for h.Len() > 0 && (*h)[0].Expiration < now {
        heap.Pop(h)
    }
}
