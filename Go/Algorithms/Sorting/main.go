package main

import "fmt"


func main() {
	Array := []int{9, 3, 2, 5, 6, 3, 4, 72}
	bubbleSort(Array)

	fmt.Print("Sorted array: ")
	for _, v := range Array {
		fmt.Print(v, " ")
	}
	fmt.Println()
}