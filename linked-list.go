package main

import "fmt"


type Node struct {
	value int
	next *Node
	prev *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Find(value int) *Node {
	node := l.head
	for node != nil {
		if node.value == value {
			return node
		} else {
			node = node.next
		}
	}
	return nil
}

func (l *List) Pop() {
	if l.tail == nil {
		return
	}
	if l.tail.prev != nil {
		l.tail = l.tail.prev
		l.tail.next = nil
	} else {
		l.head = nil
		l.tail = nil
	}
}

func (l *List) Push(value int) {
	node := &Node {value: value}

	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

func (l *List) Remove(value int) {
	node := l.Find(value)
	if node == nil {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.tail = node.prev
	}
}

func (n *Node) Next() *Node {
	return n.next
}

func LinkedList() {
	list := &List{}
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(6)
	list.Push(10)
	list.Push(5)

	node := list.First()

	for {
		println(node.value)
		node = node.Next()
		if node == nil {
			break
		}
	}

	node = list.Last()
	for {
		println(node.value)
		node = node.Prev()
		if node == nil {
			break
		}
	}

	node = list.Find(5)
	if node != nil {
		println("Found Node", node.value)
	} else {
		println("Node not found")
	}


	list.Pop()
	node = list.First()
	fmt.Println("My List after Pop: ")
	for node != nil {
		fmt.Println(node.value)
		node = node.Next()
	}

	list.Remove(6)
	node = list.First()
	fmt.Println("My List after Remove: ")
	for node != nil {
		fmt.Println(node.value)
		node = node.Next()
	}

}