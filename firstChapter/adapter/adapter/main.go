package adapter

import (
	"dsa/firstChapter/adapter/adaptee"
	"fmt"
)

type IProcess interface {
	process()
}

type Adapter struct {
	adaptee adaptee.Adaptee
}

func (adapter Adapter) process() {
	fmt.Println("Adapter Process")
	adapter.process()
}
