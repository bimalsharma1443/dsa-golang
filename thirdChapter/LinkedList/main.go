package main

import "fmt"

// node class
type Node struct {
	value int
	next  *Node
}

// linkedlist class
type LinkedList struct {
	headNode *Node
}

// AddToHead
func (list *LinkedList) AddToHead(value int) {
	var node *Node = &Node{value: value}
	if list.headNode != nil {
		node.next = list.headNode
	}
	list.headNode = node
}

// IterateList
func (list *LinkedList) IterateList() {
	var node *Node
	for node = list.headNode; node != nil; node = node.next {
		fmt.Println("node value is : ", node.value)
	}
}

// LastNode
func (list *LinkedList) LastNode() *Node {
	var node *Node
	for node = list.headNode; node != nil; node = node.next {
		if node.next == nil {
			break
		}
	}
	return node
}

// AddToEnd
func (list *LinkedList) AddToEnd(value int) {
	var node *Node = &Node{value: value}
	var endNode = list.LastNode()
	endNode.next = node
}

// NodeWithValue
func (list *LinkedList) NodeWithValue(value int) *Node {
	var node *Node
	for node = list.headNode; node != nil; node = node.next {
		if node.value == value {
			break
		}
	}
	return node
}

// AddAfter
func (list *LinkedList) AddAfter(findvalue, addValue int) {
	var oldnode *Node = list.NodeWithValue(findvalue)
	var node *Node = &Node{value: addValue}
	node.next = oldnode.next
	oldnode.next = node
}

func main() {
	var list *LinkedList = &LinkedList{}
	list.AddToHead(2)
	list.AddAfter(2, 3)
	list.AddAfter(3, 5)
	list.AddAfter(3, 4)
	list.AddToEnd(6)
	list.AddToEnd(7)
	list.IterateList()
}
