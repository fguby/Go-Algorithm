package leetcode

type LRUCache struct {
	nums     int
	capacity int
	cache    *cache
}

type cache struct {
	key   int
	value int
	node  *cache
}

func ConstructorW(capacity int) LRUCache {
	return LRUCache{
		nums:     0,
		capacity: capacity,
		cache:    nil,
	}
}

func (this *LRUCache) Get(key int) int {

	if this.nums == 0 {
		return -1
	}

	if this.nums == 1 {
		if this.cache.key == key {
			return this.cache.value
		}
		return -1
	}
	//get的节点本来就是头节点
	if this.cache.key == key {
		return this.cache.value
	}

	_, value := this.cache.get(key)

	if value == -1 {
		return value
	}

	temp := &cache{
		key:   key,
		value: value,
		node:  this.cache,
	}

	this.cache = temp

	return value

}

func (c *cache) get(key int) (*cache, int) {
	if c == nil {
		return nil, -1
	}

	if c.key == key {
		return c.node, c.value
	}

	var value int

	c.node, value = c.node.get(key)

	return c, value
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if this.cache == nil {
		this.cache = &cache{
			key:   key,
			value: value,
			node:  nil,
		}
		this.nums = 1
		return
	}
	c, ok := this.cache.put(key, value)
	if ok == -1 {
		//判断容量是否相等
		if this.nums == this.capacity {
			this.Delete()
		}
		this.nums++
	}
	head := &cache{
		key:   key,
		value: value,
		node:  c,
	}
	this.cache = head
}

func (c *cache) put(key int, value int) (*cache, int) {
	if c == nil {
		return nil, -1
	}
	if c.key == key {
		return c.node, 1
	}

	var v int
	c.node, v = c.node.put(key, value)

	return c, v
}

func (this *LRUCache) Delete() {
	if this.nums == 1 {
		this.cache = nil
		this.nums = 0
		return
	}
	this.cache = this.cache.delete()
	this.nums -= 1
}

func (c *cache) delete() *cache {
	if c.node == nil {
		return c.node
	}

	c.node = c.node.delete()

	return c
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
