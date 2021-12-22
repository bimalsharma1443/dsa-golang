package main

import "github.com/bimalsharma1443/dsa-golang/firstChapter/adapter/adapter"

func main() {

	var processor adapter.IProcess = adapter.Adapter{}
	processor.process()
}
