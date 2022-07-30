package main

import "fmt"

func div(number int) int {

	x := number / 5
	return x
}

func mult(number int) int {
	y := number * 3
	return y
}

func main() {
	var number int = 30
	mult1 := mult(18)
	div1 := div(7)
	if number < 36 {
		fmt.Printf("Number1 division five: %v\n", div1)
	} else {
		fmt.Printf("Number2 multiplied by three: %v\n", mult1)
	}
}
