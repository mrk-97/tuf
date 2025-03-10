package main

import "fmt"

func printDiagonalNumbers(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

func printDiagonalReverseStarPattern(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func main() {
	// printDiagonalNumbers(5)
	printDiagonalReverseStarPattern(5)
}
