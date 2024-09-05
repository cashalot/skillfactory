package main

type linkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (l *linkedList) Append(num int) {
	node := &Node{value: num}
	if l.length == 0 {
		l.head = node
		l.tail = node
	}
	node.prev = l.tail
	l.tail.next = node
	l.tail = node
}

func main() {
	list := &linkedList{}
	list.Append(1)
}
