package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mux sync.Mutex

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
	l.mux.Lock()
	defer l.mux.Unlock()

	if val, ok := l.items[key]; ok {
		val.Value = cacheItem{
			key:   string(key),
			value: value,
		}
		l.queue.MoveToFront(val)
		return true
	}

	if l.queue.Len()+1 > l.capacity {
		back := l.queue.Back()
		l.queue.Remove(back)
		delete(l.items, Key(back.Value.(cacheItem).key))
	}
	newListItem := l.queue.PushFront(cacheItem{
		key:   string(key),
		value: value,
	})
	l.items[key] = newListItem

	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mux.Lock()
	defer l.mux.Unlock()

	val, existsInMap := l.items[key]
	if l.queue.Len() == 0 || !existsInMap {
		return nil, false
	}

	l.queue.MoveToFront(val)
	return val.Value.(cacheItem).value, true
}

func (l *lruCache) Clear() {
	l.mux.Lock()
	defer l.mux.Unlock()

	l.queue = NewList()
}
