package main

import (
	"fmt"
	"os"
)

func isEvenOrOdd(num int) {
	if num%2 == 0 {
		fmt.Printf("Even\n")
	} else {
		fmt.Printf("Odd\n")
	}
}

func gradeAssessment(score int) {
	switch {
	case (score >= 90 && score <= 100):
		fmt.Println("A")
	case (score >= 80 && score <= 89):
		fmt.Println("B")
	case (score >= 70 && score <= 79):
		fmt.Println("C")
	case (score >= 60 && score <= 69):
		fmt.Println("D")
	case (score < 60):
		fmt.Println("F")
	default:
		fmt.Println("Invalid score")
	}
}

func removeWhitespace(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			return s[:i] + s[i+1:]
		}
	}
	return s
}

func finxMax(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func sumOfSquares(n int) int {
	if n <= 0 {
		fmt.Print("Eneter non negative number")
		os.Exit(1)
	}
	sum := 0
	for i := 1; i <= n; i++ {
		squares := i * i
		sum += squares
	}
	return sum

}

func countDown(end int) {
	start := 0
	for end > start {
		fmt.Println(end)
		end--
		fmt.Print("Countdown Completed\n")
	}
}

func calculateBMI(weight float32, height float32) float32 {
	BMI := weight / (height * height)
	return BMI
}

func main() {
	isEvenOrOdd(27)
	gradeAssessment(93)
	fmt.Println(removeWhitespace("Hello World"))
	max := finxMax([]int{1, 4, 4, 19})
	fmt.Printf("Max val of array: %v\n", max)

	fmt.Println("sum of squares:", sumOfSquares(4))
	countDown(3)

	BMI := calculateBMI(91.8, 167.00)
	fmt.Println("BMI:", BMI)
}
