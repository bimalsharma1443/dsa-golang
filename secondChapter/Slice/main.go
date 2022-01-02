package main

import "fmt"

func main() {
	sliceData := []int{1, 2, 3, 4, 5}
	fmt.Println("Capacity of sliceData : ", cap(sliceData))
	fmt.Println("len of SliceData : ", len(sliceData))

	sliceData = append(sliceData, 8)
	fmt.Println("Capacity of sliceData : ", cap(sliceData))
	fmt.Println("len of SliceData : ", len(sliceData))

	twoSliceData := make([][]int, 7)
	for i := range twoSliceData {
		twoSliceData[i] = make([]int, 7)
	}

	fmt.Println(twoSliceData)
}
