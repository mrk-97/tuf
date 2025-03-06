package main

import "fmt"

func reverseArray(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}

func main() {
	fmt.Println(reverseArray([]int{1, 2, 3, 4, 5}))
}
