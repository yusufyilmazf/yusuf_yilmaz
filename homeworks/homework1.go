package main

import "fmt"

func div(number int) (x int) {
	x = number / 5
	return
}

func mult(number int) (y int) {
	y = number * 3
	return
}

func main() {
	fmt.Printf("Number division five: %d\n", div(25))
	fmt.Printf("Number multiplied by three: %d\n", mult(25))

}
