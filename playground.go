package main

import "fmt"

func multipleReturns(x int) (int, int) {
	var doubled int
	doubled = x * 2
	return x, doubled
}

func main() {
	var value, doubled int
	value, doubled = multipleReturns(2)
	fmt.Print(value, doubled)
}
