package main

import "fmt"

func contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

func reverse(input string) (output string) {
	for _, v := range input {
		output = string(v) + output
	}
	return
}

func zero(x int) {
	x = 0
}

func zeroWithPointer(x *int) {
	*x = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5
	zeroWithPointer(&x)
	fmt.Println(x) // x is now 0 as we modify the memory
	fmt.Println(&x)
}

// func main() {
// 	joseph := "one cool dude!"
// 	josephAddress := &joseph //points to the joseph memory
// 	fmt.Println(josephAddress)
// 	fmt.Println(*josephAddress)
// 	*josephAddress = "one cool and updated dude!" //updates the joseph memory
// 	fmt.Println(josephAddress)
// 	fmt.Println(*josephAddress) //both are now updated as they point to the same place
// 	fmt.Println(joseph) //both are now updated as they point to the same place
// }
