package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mutex    *sync.Mutex
}

type cacheItem struct {
	k Key
	v interface{}
}

func NewCache(capacity int) Cache {
	return &cache{
		capacity: capacity,
		items:    make(map[Key]*ListItem, capacity),
		queue:    NewList(),
		mutex:    new(sync.Mutex),
	}
}

func (c *cache) Set(k Key, v interface{}) bool {
	c.mutex.Lock()

	defer c.mutex.Unlock()

	item := &cacheItem{k: k, v: v}
	if i, ok := c.items[k]; ok {
		i.Value = item
		c.queue.MoveToFront(i)

		return true
	}

	if c.queue.Len() == c.capacity {
		lastItem, ok := c.queue.Back().Value.(*cacheItem)
		if !ok {
			return false
		}

		c.queue.Remove(c.queue.Back())
		delete(c.items, lastItem.k)
	}

	pushedToFront := c.queue.PushFront(item)
	c.items[k] = pushedToFront

	return false
}

func (c *cache) Get(key Key) (interface{}, bool) {
	c.mutex.Lock()

	defer c.mutex.Unlock()

	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)

		cachedItem, ok := item.Value.(*cacheItem)
		if !ok {
			return nil, false
		}

		return cachedItem.v, true
	}

	return nil, false
}

func (c *cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
