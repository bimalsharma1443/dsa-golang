package main

import "github.com/bimalsharma1443/dsa-golang/adapter"

func main() {
	var processor adapter.IProcess = adapter.Adapter{}
	processor.process()
}
