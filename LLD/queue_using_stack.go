package LLD

import "fmt"

type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{
		in:  []int{},
		out: []int{},
	}
}

// Push element into queue
func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

// Remove front element
func (q *MyQueue) Pop() int {

	// Transfer only if out stack is empty
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			n := len(q.in)

			val := q.in[n-1]
			q.in = q.in[:n-1]

			q.out = append(q.out, val)
		}
	}

	n := len(q.out)
	val := q.out[n-1]
	q.out = q.out[:n-1]

	return val
}

// Return front element
func (q *MyQueue) Peek() int {

	// Transfer only if out stack is empty
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			n := len(q.in)

			val := q.in[n-1]
			q.in = q.in[:n-1]

			q.out = append(q.out, val)
		}
	}

	return q.out[len(q.out)-1]
}

// Check if queue is empty
func (q *MyQueue) Empty() bool {
	return len(q.in) == 0 && len(q.out) == 0
}

func main() {

	q := Constructor()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Peek()) // 1
	fmt.Println(q.Pop())  // 1
	fmt.Println(q.Pop())  // 2

	q.Push(4)

	fmt.Println(q.Pop())  // 3
	fmt.Println(q.Peek()) // 4
	fmt.Println(q.Pop())  // 4

	fmt.Println(q.Empty()) // true
}