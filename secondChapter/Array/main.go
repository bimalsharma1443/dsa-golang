package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 4, 5, 6}
	var i int
	for i = 0; i < len(arr); i++ {
		fmt.Println("Printing an element ", arr[i])
	}
	for i, val := range arr {
		fmt.Printf("index %d range %d \n", i, val)
	}
	for _, val := range arr {
		fmt.Println("without range", val)
	}
}
