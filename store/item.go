package store

type Item struct {
    Value      interface{} 
    Expiration int64       
    Index      int         
}

func NewItem(value interface{}, ttl int64) *Item {
    return &Item{
        Value:      value,
        Expiration: ttl,
        Index:      -1, 
    }
}

func (item *Item) IsExpired() bool {
    if item.Expiration == 0 {
        return false 
    }
    return item.Expiration < time.Now().UnixNano()
}
