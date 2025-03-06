package main

import "fmt"

func trafficLight(color int) string {
	switch color {
	case 1:
		return "stop"
	case 2:
		return "caution"
	case 3:
		return "go"
	default:
		return "invalid color"
	}
}

func main() {
	fmt.Println(trafficLight(2))
}
