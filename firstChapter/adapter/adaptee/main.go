package adaptee

import "fmt"

type Adaptee struct {
	adapterType int
}

func (Adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}
