package adapter

import (
	"fmt"

	"github.com/bimalsharma1443/dsa-golang/firstChapter/adapter/adaptee"
)

type IProcess interface {
	Process()
}

type Adapter struct {
	Adaptee adaptee.Adaptee
}

func (adapter Adapter) Process() {
	fmt.Println("Adapter Process")
	adapter.Adaptee.Convert()
}

func main() {
	var processor IProcess = Adapter{}
	processor.Process()
}
