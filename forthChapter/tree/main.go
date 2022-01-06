package main

import (
	"fmt"
	"sync"
)

// node class
type TreeNode struct {
	key       int
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

// Binary search tree
type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

func (binarySearchTree *BinarySearchTree) InsertElement(key, value int) {
	binarySearchTree.lock.Lock()
	defer binarySearchTree.lock.Unlock()
	var treeNode *TreeNode = &TreeNode{key: key, value: value}
	if binarySearchTree.rootNode == nil {
		binarySearchTree.rootNode = treeNode
	} else {
		insertTreeNode(binarySearchTree.rootNode, treeNode)
	}
}

func insertTreeNode(rootNode, treeNode *TreeNode) {
	if treeNode.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = treeNode
		} else {
			insertTreeNode(rootNode.leftNode, treeNode)
		}
	} else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = treeNode
		} else {
			insertTreeNode(rootNode.rightNode, treeNode)
		}
	}

}

func (binarySearchTree *BinarySearchTree) InOrderTraverseTree(function func(int)) {
	binarySearchTree.lock.Lock()
	defer binarySearchTree.lock.Unlock()
	fmt.Println("------------------------------------------------")
	inOrderTraverseTree(binarySearchTree.rootNode, function)
}

func inOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		inOrderTraverseTree(treeNode.leftNode, function)
		function(treeNode.value)
		inOrderTraverseTree(treeNode.rightNode, function)
	}
}

func (binarySearchTree *BinarySearchTree) PreOrderTraverseTree(function func(int)) {
	binarySearchTree.lock.Lock()
	defer binarySearchTree.lock.Unlock()
	fmt.Println("------------------------------------------------")
	preOrderTraverseTree(binarySearchTree.rootNode, function)
}

func preOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		function(treeNode.value)
		preOrderTraverseTree(treeNode.leftNode, function)
		preOrderTraverseTree(treeNode.rightNode, function)
	}
}

func (binarySearchTree *BinarySearchTree) PostOrderTraverseTree(function func(int)) {
	binarySearchTree.lock.Lock()
	defer binarySearchTree.lock.Unlock()
	fmt.Println("------------------------------------------------")
	postOrderTraverseTree(binarySearchTree.rootNode, function)
}

func postOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		postOrderTraverseTree(treeNode.leftNode, function)
		postOrderTraverseTree(treeNode.rightNode, function)
		function(treeNode.value)
	}
}

func (binarySearchTree *BinarySearchTree) MinNode() *int {
	binarySearchTree.lock.RLock()
	defer binarySearchTree.lock.RUnlock()
	var treeNode *TreeNode
	treeNode = binarySearchTree.rootNode
	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.leftNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.leftNode
	}
}

func (binarySearchTree *BinarySearchTree) MaxNode() *int {
	binarySearchTree.lock.RLock()
	defer binarySearchTree.lock.RUnlock()
	var treeNode *TreeNode
	treeNode = binarySearchTree.rootNode
	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.rightNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.rightNode
	}
}

func (binarySearchTree *BinarySearchTree) Search(key int) bool {
	binarySearchTree.lock.RLock()
	defer binarySearchTree.lock.RUnlock()
	return search(binarySearchTree.rootNode, key)
}

func search(treenode *TreeNode, key int) bool {
	if treenode == nil {
		return false
	}

	if treenode.key > key {
		return search(treenode.leftNode, key)
	}

	if treenode.key < key {
		return search(treenode.rightNode, key)
	}

	return true
}

func (binarySearchTree *BinarySearchTree) RemoveNode(key int) {
	binarySearchTree.lock.Lock()
	defer binarySearchTree.lock.Unlock()
	removeNode(binarySearchTree.rootNode, key)
}

func removeNode(treeNode *TreeNode, key int) *TreeNode {
	if treeNode == nil {
		return nil
	}
	if key < treeNode.key {
		treeNode.leftNode = removeNode(treeNode.leftNode, key)
		return treeNode
	}
	if key > treeNode.key {
		treeNode.rightNode = removeNode(treeNode.rightNode, key)
		return treeNode
	}
	// key == node.key
	if treeNode.leftNode == nil && treeNode.rightNode == nil {
		treeNode = nil
		return nil
	}
	if treeNode.leftNode == nil {
		treeNode = treeNode.rightNode
		return treeNode
	}
	if treeNode.rightNode == nil {
		treeNode = treeNode.leftNode
		return treeNode
	}
	var leftmostrightNode *TreeNode
	leftmostrightNode = treeNode.rightNode
	for {
		//find smallest value on the right side
		if leftmostrightNode != nil && leftmostrightNode.leftNode != nil {
			leftmostrightNode = leftmostrightNode.leftNode
		} else {
			break
		}
	}
	treeNode.key, treeNode.value = leftmostrightNode.key, leftmostrightNode.value
	treeNode.rightNode = removeNode(treeNode.rightNode, treeNode.key)
	return treeNode
}

func (binarySearchTree *BinarySearchTree) String() {
	binarySearchTree.lock.Lock()
	defer binarySearchTree.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(binarySearchTree.rootNode, 0)
	fmt.Println("------------------------------------------------")

}

func stringify(treeNode *TreeNode, level int) {
	if treeNode != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += " "
		}
		format += "---[ "
		level++
		stringify(treeNode.leftNode, level)
		fmt.Printf(format+"%d\n", treeNode.key)
		stringify(treeNode.rightNode, level)
	}
}

func main() {
	var binarySearchTree *BinarySearchTree = &BinarySearchTree{}
	binarySearchTree.InsertElement(15, 15)
	binarySearchTree.InsertElement(5, 5)
	binarySearchTree.InsertElement(20, 20)
	binarySearchTree.InsertElement(3, 3)
	binarySearchTree.InsertElement(8, 8)
	binarySearchTree.InsertElement(16, 25)
	binarySearchTree.InsertElement(7, 7)
	binarySearchTree.InsertElement(11, 11)
	binarySearchTree.InsertElement(10, 10)
	binarySearchTree.InOrderTraverseTree(func(i int) { fmt.Println(i) })
	// binarySearchTree.String()
	// binarySearchTree.PreOrderTraverseTree(func(i int) { fmt.Println(i) })
	// binarySearchTree.PostOrderTraverseTree(func(i int) { fmt.Println(i) })
	// fmt.Println("Min node is : ", *binarySearchTree.MinNode())
	// fmt.Println("Max node is : ", *binarySearchTree.MaxNode())
	// if binarySearchTree.Search(10) {
	// 	fmt.Println("node is available")
	// } else {
	// 	fmt.Println("node is not available")
	// }
	binarySearchTree.RemoveNode(8)
	binarySearchTree.InOrderTraverseTree(func(i int) { fmt.Println(i) })
}
