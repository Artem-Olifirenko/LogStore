package store

import (
    "container/heap"
    "sync"
    "time"
)

type Item struct {
    Key        string
    Value      interface{}
    Expiration int64
    Index      int 
}

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
    *h = old[:n-1]
    return item
}

type Store struct {
    mu    sync.RWMutex
    items map[string]*Item
    pq    *ItemHeap 
}

func NewStore() *Store {
    pq := &ItemHeap{}
    heap.Init(pq)
    store := &Store{
        items: make(map[string]*Item),
        pq:    pq,
    }
    
    go store.cleanupExpiredItemsRoutine()
    return store
}

func (s *Store) Set(key string, value interface{}, ttl time.Duration) {
    s.mu.Lock()
    defer s.mu.Unlock()

    exp := int64(0)
    if ttl > 0 {
        exp = time.Now().Add(ttl).UnixNano()
    }

    if item, exists := s.items[key]; exists {
     
        item.Value = value
        item.Expiration = exp
        heap.Fix(s.pq, item.Index)
    } else {
       
        item := &Item{
            Key:        key,
            Value:      value,
            Expiration: exp,
        }
        s.items[key] = item
        heap.Push(s.pq, item) 
    }
}

func (s *Store) Get(key string) (interface{}, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    item, exists := s.items[key]
    if !exists || (item.Expiration > 0 && time.Now().UnixNano() > item.Expiration) {
        return nil, false
    }

    return item.Value, true
}

func (s *Store) Delete(key string) {
    s.mu.Lock()
    defer s.mu.Unlock()

    if item, exists := s.items[key]; exists {
        heap.Remove(s.pq, item.Index) 
        delete(s.items, key)
    }
}

func (s *Store) cleanupExpiredItemsRoutine() {
    ticker := time.NewTicker(1 * time.Minute)
    for {
        <-ticker.C
