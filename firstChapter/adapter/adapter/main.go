package adapter

import (
	"fmt"

	"github.com/bimalsharma1443/dsa-golang/firstChapter/adapter/adaptee"
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
