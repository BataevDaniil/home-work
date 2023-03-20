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

func (c *lruCache) Set(key Key, v interface{}) bool {
	item := c.items[key]
	if item == nil {
		c.queue.PushFront(v)
		if c.queue.Len() > c.capacity {
			c.queue.Remove(c.queue.Back())
			delete(c.items, key)
		}
		c.items[key] = c.queue.Front()
		return false
	} else {
		c.queue.MoveToFront(item)
		item.Value = v
		return true
	}
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	item := c.items[key]
	if item == nil {
		return nil, false
	} else {
		return item.Value, true
	}
}

func (c *lruCache) Clear() {
	for ; c.queue.Len() > 0; c.queue.Remove(c.queue.Front()) {
	}
	for k := range c.items {
		delete(c.items, k)
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
