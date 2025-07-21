package kvmap

type entry struct {
	key   string
	value int
	prev  *entry
	next  *entry
}

type LRUCache struct {
	capacity int
	items    map[string]*entry
	head     *entry
	tail     *entry
}

func NewLRU(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*entry),
	}
}

func (c *LRUCache) Get(key string) (int, bool) {
	if node, exists := c.items[key]; exists {
		c.moveToFront(node)
		return node.value, true
	}
	return 0, false
}

func (c *LRUCache) Put(key string, value int) {
	if node, exists := c.items[key]; exists {
		node.value = value
		c.moveToFront(node)
		return
	}

	newEntry := &entry{key: key, value: value}
	c.items[key] = newEntry
	c.addToFront(newEntry)

	if len(c.items) > c.capacity {
		evicted := c.removeTail()
		delete(c.items, evicted.key)
	}
}

func (c *LRUCache) moveToFront(e *entry) {
	if c.head == e {
		return
	}
	c.remove(e)
	c.addToFront(e)
}

func (c *LRUCache) addToFront(e *entry) {
	e.next = c.head
	e.prev = nil
	if c.head != nil {
		c.head.prev = e
	}
	c.head = e
	if c.tail == nil {
		c.tail = e
	}
}

func (c *LRUCache) remove(e *entry) {
	if e.prev != nil {
		e.prev.next = e.next
	} else {
		c.head = e.next
	}
	if e.next != nil {
		e.next.prev = e.prev
	} else {
		c.tail = e.prev
	}
}

func (c *LRUCache) removeTail() *entry {
	if c.tail == nil {
		return nil
	}
	tail := c.tail
	c.remove(tail)
	return tail
}
