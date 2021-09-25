package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

// type cacheItem struct {
//	key   string
//	value interface{}
// }

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	wasInMap := false
	if val, ok := l.items[key]; ok {
		l.queue.Remove(val)
		wasInMap = true
	}

	newListItem := l.queue.PushFront(value)
	l.items[key] = newListItem
	if !wasInMap && l.queue.Len() > l.capacity {
		l.queue.Remove(l.queue.Back())
		delete(l.items, key)
	}
	return wasInMap
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	val, existsInMap := l.items[key]
	if l.queue.Len() == 0 || !existsInMap {
		return nil, false
	}

	l.queue.MoveToFront(val)
	return val.Value, true
}

func (l *lruCache) Clear() {
	l.queue = NewList()
}
