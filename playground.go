package main

import (
	"fmt"
)

func averageArray(arr [4]int) int {
	var total int

	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	return total
}

type Salary struct {
	basic, allowance, bonuses int
}

type Employee struct {
	firstName, lastName string
	salary              Salary
	fullTime            bool
}

func mapStuff() {
	testMap := map[string]int{
		"firstest": 1,
		"second":   2,
		"third":    3,
		"4th":      4,
	}
	for key, value := range testMap {
		fmt.Println(key, "=>", value)
	}
}

func main() {
	joe := Employee{
		firstName: "joe",
		lastName:  "phwllng",
		salary:    Salary{10, 20, 30},
		fullTime:  false,
	}
	fmt.Println(joe.salary.basic)
	joe.salary.basic = 20
	fmt.Println(joe.salary.basic)
	mapStuff()
}
