package main

import "fmt"

type stack struct {
	items []int
}

func (s *stack) Push(value int)  {
	s.items = append(s.items, value)
}

func (s *stack) Pop() int {
	stackLenn := len(s.items)
	if stackLenn == 0 {
		return -1
	}
	item, items := s.items[stackLenn - 1], s.items[0:stackLenn - 1]
	s.items = items
	return item
}

func Stack() {

	stack := stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	fmt.Println("my stack is: ", stack)

}