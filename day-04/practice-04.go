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

func main() {
	printStarPattern(5)
}
