package main

import (
	"bufio"
	"fmt"
	"os"
)

func UserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	output := UserInput()

	fmt.Println(output)

}
