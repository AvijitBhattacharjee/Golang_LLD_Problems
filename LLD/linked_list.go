package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Insert at beginning
func (l *LinkedList) InsertFront(val int) {
	node := &Node{data: val}
	node.next = l.head
	l.head = node
}

// Insert at end
func (l *LinkedList) InsertEnd(val int) {

	node := &Node{data: val}

	if l.head == nil {
		l.head = node
		return
	}

	curr := l.head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = node
}

// Insert at position (0-based)
func (l *LinkedList) InsertAt(pos, val int) {

	if pos == 0 {
		l.InsertFront(val)
		return
	}

	curr := l.head

	for i := 0; i < pos-1 && curr != nil; i++ {
		curr = curr.next
	}

	if curr == nil {
		return
	}

	node := &Node{data: val}

	node.next = curr.next
	curr.next = node
}

// Delete front
func (l *LinkedList) DeleteFront() {

	if l.head == nil {
		return
	}

	l.head = l.head.next
}

// Delete end
func (l *LinkedList) DeleteEnd() {

	if l.head == nil {
		return
	}

	if l.head.next == nil {
		l.head = nil
		return
	}

	curr := l.head

	for curr.next.next != nil {
		curr = curr.next
	}

	curr.next = nil
}

// Delete by value
func (l *LinkedList) DeleteValue(val int) {

	if l.head == nil {
		return
	}

	if l.head.data == val {
		l.head = l.head.next
		return
	}

	curr := l.head

	for curr.next != nil && curr.next.data != val {
		curr = curr.next
	}

	if curr.next != nil {
		curr.next = curr.next.next
	}
}

// Search
func (l *LinkedList) Search(val int) bool {

	curr := l.head

	for curr != nil {

		if curr.data == val {
			return true
		}

		curr = curr.next
	}

	return false
}

// Reverse
func (l *LinkedList) Reverse() {

	var prev *Node
	curr := l.head

	for curr != nil {

		next := curr.next

		curr.next = prev

		prev = curr

		curr = next
	}

	l.head = prev
}

// Length
func (l *LinkedList) Length() int {

	count := 0

	curr := l.head

	for curr != nil {

		count++

		curr = curr.next
	}

	return count
}

// Print
func (l *LinkedList) Print() {

	curr := l.head

	for curr != nil {

		fmt.Print(curr.data)

		if curr.next != nil {
			fmt.Print(" -> ")
		}

		curr = curr.next
	}

	fmt.Println()
}

func main() {

	list := &LinkedList{}

	list.InsertEnd(10)
	list.InsertEnd(20)
	list.InsertEnd(30)

	list.Print()

	list.InsertFront(5)

	list.Print()

	list.InsertAt(2, 15)

	list.Print()

	list.DeleteFront()

	list.Print()

	list.DeleteEnd()

	list.Print()

	list.DeleteValue(15)

	list.Print()

	fmt.Println("Length:", list.Length())

	fmt.Println("Search 20:", list.Search(20))
	fmt.Println("Search 100:", list.Search(100))

	list.Reverse()

	list.Print()
}