package main

import "fmt"

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {

	// enqueue
	s.queue = append(s.queue, x)

	// rotate previous elements
	n := len(s.queue)
	for i := 0; i < n-1; i++ {

		front := s.queue[0]
		s.queue = s.queue[1:]
		s.queue = append(s.queue, front)
	}
}

func (s *MyStack) Pop() int {

	top := s.queue[0]
	s.queue = s.queue[1:]
	return top
}

func (s *MyStack) Top() int {
	return s.queue[0]
}

func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}

func main() {

	s := Constructor()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Pop()) // 3
	fmt.Println(s.Top()) // 2
	fmt.Println(s.Pop()) // 2
	fmt.Println(s.Empty())
}