package main

import "dsa/firstChapter/adapter/adapter"

func main() {
	var processor adapter.IProcess = adapter.Adapter{}

	processor.process()
}
