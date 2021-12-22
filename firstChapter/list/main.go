package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List

	intList.PushBack(11)
	intList.PushBack(12)
	intList.PushBack(13)

	for ele := intList.Front(); ele != nil; ele = ele.Next() {
		fmt.Println("value is :", ele.Value)
	}
}
