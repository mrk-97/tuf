package main

import "fmt"

func printStarPattern(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func printDiagonalPattern(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func printDiagonalNumberPattern(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			fmt.Print(j)
		}
		fmt.Println(i)
	}
}
func printDiagonalNumberPattern11(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			fmt.Print(j)
		}
		fmt.Println(i + i)
	}
}

func main() {
	printStarPattern(5)
	printDiagonalPattern(5)
	printDiagonalNumberPattern(5)
	printDiagonalNumberPattern11(4)

}
