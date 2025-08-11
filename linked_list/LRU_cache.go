package linkedlist

type DoubleNode struct {
	Key  int
	Val  int
	Next *DoubleNode
	Prev *DoubleNode
}

type LRUCache struct {
	capacity int
	head     *DoubleNode
	tail     *DoubleNode
	cache    map[int]*DoubleNode
}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DoubleNode),
	}
}

func (c *LRUCache) moveToHead(node *DoubleNode) {
	if node == c.head {
		return // already MRU
	}

	// detach node
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	// update tail if needed
	if node == c.tail {
		c.tail = node.Prev
	}

	// insert node at head
	node.Prev = nil
	node.Next = c.head
	if c.head != nil {
		c.head.Prev = node
	}
	c.head = node

	// if list was empty, set tail
	if c.tail == nil {
		c.tail = node
	}
}

// Remove the tail node (LRU)
func (c *LRUCache) removeTail() {
	if c.tail == nil {
		return
	}
	delete(c.cache, c.tail.Key)
	if c.tail.Prev != nil {
		c.tail.Prev.Next = nil
	} else {
		c.head = nil // cache becomes empty
	}
	c.tail = c.tail.Prev
}

func (c *LRUCache) Get(key int) int {
	if node, exists := c.cache[key]; exists {
		c.moveToHead(node)
		return node.Val
	}
	return -1
}

// Put key-value pair
func (c *LRUCache) Put(key int, value int) {
	if node, exists := c.cache[key]; exists {
		node.Val = value
		c.moveToHead(node)
		return
	}

	newNode := &DoubleNode{Key: key, Val: value}
	c.cache[key] = newNode
	c.moveToHead(newNode)

	if len(c.cache) > c.capacity {
		c.removeTail()
	}
}
