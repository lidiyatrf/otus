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

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if val, ok := l.items[key]; ok {
		val.Value = cacheItem{
			key:   string(key),
			value: value,
		}
		l.queue.MoveToFront(val)
		return true
	}

	newListItem := l.queue.PushFront(cacheItem{
		key:   string(key),
		value: value,
	})
	l.items[key] = newListItem
	if l.queue.Len() > l.capacity {
		l.queue.Remove(l.queue.Back())
		delete(l.items, key)
	}
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	val, existsInMap := l.items[key]
	if l.queue.Len() == 0 || !existsInMap {
		return nil, false
	}

	l.queue.MoveToFront(val)
	return val.Value.(cacheItem).value, true
}

func (l *lruCache) Clear() {
	l.queue = NewList()
}
