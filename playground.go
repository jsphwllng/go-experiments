package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
	const test string = "hello"
	s := strings.Split(test, "")
	for i, s := range s {
		fmt.Println(i, s)
	}
}

//! these are not syncronised
func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}
