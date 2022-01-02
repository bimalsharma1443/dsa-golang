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
	binarySearchTree.InsertElement(8, 8)
	binarySearchTree.InsertElement(3, 3)
	binarySearchTree.InsertElement(10, 10)
	binarySearchTree.InsertElement(1, 1)
	binarySearchTree.InsertElement(6, 6)
	binarySearchTree.String()
}
