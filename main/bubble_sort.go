package main

import (
	"fmt"
)

var arr = []int{10, 55, 22, 41, 65, 85, 12}

func bubblesort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}

func main() {
	fmt.Println("1", arr)
	fmt.Println("2", bubblesort(arr))
}
