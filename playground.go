package main

import (
	"fmt"
	"strings"
)

func main() {
	salutation := "Hello!"
	name := "joe"
	sayGreeting(&salutation, &name)
	fmt.Println(salutation, name)
	scream("hello", "goodbye", "nihao", "mabrook", "oh")
}

func sayGreeting(salutation, name *string) {
	*salutation = "go away"
	*name = "michael"
}

func scream(words ...string) {
	for _, v := range words {
		fmt.Println(strings.ToUpper(v) + "!!!!!!!!!!!!!")
	}
}
