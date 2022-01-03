package main

import "fmt"

// create node class
type Node struct {
	value int
	next  *Node
	prev  *Node
}

// create link list class
type LinkedList struct {
	headNode *Node
}

// NodeBetweenValues
func (list *LinkedList) NodeBetweenValues(prevValue, nextValue int) (node *Node) {
	for node = list.headNode; node != nil; node = node.next {
		if node.next != nil && node.prev != nil {
			if node.prev.value == prevValue && node.next.value == nextValue {
				break
			}
		}
	}
	return node
}

// AddToHead
func (list *LinkedList) AddToHead(value int) {
	var node *Node = &Node{value: value}
	if list.headNode != nil {
		node.next = list.headNode
		list.headNode.prev = node
	}

	list.headNode = node
}

// IterateList
func (list *LinkedList) IterateList() {
	var node *Node
	for node = list.headNode; node != nil; node = node.next {
		fmt.Printf("Prev Node : %v, Node : %d, Next Node %v \n", checkPrevValue(node), node.value, checkNextValue(node))
	}
}

func checkPrevValue(node *Node) int {
	if node.prev != nil {
		return node.prev.value
	}
	return 0
}

func checkNextValue(node *Node) int {
	if node.next != nil {
		return node.next.value
	}
	return 0
}

// AddAfter
func (list *LinkedList) AddAfter(prevValue, newValue int) {
	var node *Node = &Node{value: newValue}
	var oldNode *Node = list.NodeWithValue(prevValue)
	fmt.Println(oldNode)
	node.next = oldNode.next
	node.next.prev = node
	oldNode.next = node
	node.prev = oldNode

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

// AddToEnd
func (list *LinkedList) AddToEnd(value int) {
	var node *Node = &Node{value: value}
	var endNode = list.LastNode()
	endNode.next = node
	node.prev = endNode
}

func (list *LinkedList) LastNode() *Node {
	var node *Node
	for node = list.headNode; node != nil; node = node.next {
		if node.next == nil {
			break
		}
	}
	return node
}

// NodeWithPrevNode
func (list *LinkedList) NodeWithPrevNode(value int) *Node {
	var node *Node
	for node = list.headNode; node != nil; node = node.next {
		if node.next != nil && node.next.value == value {
			break
		}
	}
	return node
}

// DeleteLinklist
func (list *LinkedList) DeleteLinklist(value int) {
	var currentNode *Node = list.NodeWithValue(value)
	currentNode.prev.next = currentNode.next
	currentNode.next.prev = currentNode.prev
}

func main() {
	var list *LinkedList = &LinkedList{}
	list.AddToHead(3)
	list.AddToHead(2)
	list.AddToEnd(5)
	list.AddAfter(3, 4)
	list.AddAfter(3, 10)
	list.AddToEnd(6)
	list.IterateList()
	list.DeleteLinklist(10)
	list.IterateList()
	node := list.NodeBetweenValues(3, 5)

	if node != nil {
		fmt.Println("node value : ", node.value)
	}

}
