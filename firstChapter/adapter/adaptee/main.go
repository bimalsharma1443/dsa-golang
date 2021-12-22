package adaptee

import "fmt"

type Adaptee struct {
	adapterType int
}

func (Adaptee Adaptee) Convert() {
	fmt.Println("Adaptee convert method")
}
