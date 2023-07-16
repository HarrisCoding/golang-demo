package lru

type Node struct {
	key, val string
	prev     *Node
	next     *Node
}

// NodeList 为双向链表, 新加入 Node 放到 tail
type NodeList struct {
	head *Node
	tail *Node
}

func (l *NodeList) addNode(key, val string) *Node {
	newNode := &Node{
		prev: l.tail.prev,
		next: l.tail,
		key:  key,
		val:  val,
	}
	l.tail.prev.next = newNode
	l.tail.prev = newNode
	return newNode
}

// makeRecently 将 node 移动到链表尾部
func (l *NodeList) makeRecently(node *Node) {
	l.deleteNode(node)
	l.tail.prev.next = node
	node.prev = l.tail.prev
	l.tail.prev = node
	node.next = l.tail
}

// deleteNode 删除链表当前节点
func (l *NodeList) deleteNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// deleteLeastRecently 淘汰最久未用的节点（链表头部）
func (l *NodeList) deleteLeastRecently() *Node {
	if l.head.next == l.tail {
		return nil
	}
	first := l.head.next
	l.deleteNode(l.head.next)
	return first
}

type LRUCache struct {
	capacity int
	list     NodeList
	nodeMap  map[string]*Node
}

func NewLRUCache(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head
	list := NodeList{
		head: head,
		tail: tail,
	}
	return LRUCache{
		capacity: capacity,
		nodeMap:  make(map[string]*Node),
		list:     list,
	}
}

// Get 获取 key 对应值
func (c LRUCache) Get(key string) string {
	if node, ok := c.nodeMap[key]; ok {
		c.list.makeRecently(node)
		return node.val
	}
	return ""
}

// Put 写入 key 对应值
func (c LRUCache) Put(key, val string) {
	if c.capacity == 0 {
		return
	}
	if node, ok := c.nodeMap[key]; ok {
		node.val = val
		c.list.makeRecently(node)
		return
	}

	// 达到最大容量
	if c.capacity == len(c.nodeMap) {
		node := c.list.deleteLeastRecently()
		if node != nil {
			delete(c.nodeMap, node.key)
		}
	}

	newNode := c.list.addNode(key, val)
	c.nodeMap[key] = newNode
}
