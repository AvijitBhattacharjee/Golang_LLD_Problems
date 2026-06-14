package LLD

import "fmt"

type Node struct {
	key int
	value int
	prev *Node
	next *Node
}

type LRUCache struct {
	head *Node
	tail *Node
	data map[int]*Node
	capacity int
}

func New_Cache(capacity int) *LRUCache {

	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		head: head,
		tail: tail,
		data: make(map[int]*Node),
	}
}

func (this *LRUCache)Get(key int) int {

	node, ok := this.data[key]
	if !ok {
		return -1
	}
	this.remove(node)
	this.insertAfterHead(node)

	return node.value
}

func (this *LRUCache)Set(key, value int) {
	

	if node, ok := this.data[key]; ok {

		node.value = value
		this.remove(node)
		this.insertAfterHead(node)
		return
	}

	var newNode = &Node{
		key: key,
		value: value,
	}

	this.data[key] = newNode
	this.insertAfterHead(newNode)

	if len(this.data) > this.capacity {
		this.remove(this.tail.prev)
		delete(this.data, this.tail.prev.key)
	}

}

func (l *LRUCache)remove(node *Node) {

	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
}

func (l *LRUCache)insertAfterHead(node *Node) {

	node.prev = l.head
	node.next = l.head.next

	l.head.next.prev = node
	l.head.next = node
}


func main() {
	
	var l = New_Cache(2)

	l.Set(1,1)
	l.Set(1,1)

	fmt.Println(l.Get(1))

	l.Set(3,3)

	fmt.Println(l.Get(2))
	fmt.Println(l.Get(3))
}