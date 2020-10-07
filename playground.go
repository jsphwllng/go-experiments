package main

import (
	"fmt"
	"sync"
	"time"
)

//wg = "wait group"
var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("recovered in cleanup: ", r)
	}
}

func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 0 {
			panic("an example error that stops the script.")
		}
	}
}

func main() {
	defer fmt.Println("this is a defer")
	wg.Add(1)
	go say("hello!")
	wg.Add(1)
	go say("neato")
	wg.Wait()
}
