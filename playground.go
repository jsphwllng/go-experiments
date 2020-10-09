package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//seeding
	rand.Seed(time.Now().UnixNano())

	c1, c2 := make(chan string, 1), make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second): //this will wait for one second
		fmt.Println("timeout 1")
	}
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second): //this will wait for three seconds
		fmt.Println("timeout 2")
	}

	//we could also do it this way

	c3 := make(chan string, 1)
	n := rand.Intn(5)
	fmt.Println("who painted the mona lisa?")
	timeout := time.After(time.Duration(n) * time.Second)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(5 * time.Second)
			c3 <- "da vinky?"
		}
	}()
	select {
	case res := <-c3:
		fmt.Println(res)
	case <-timeout:
		fmt.Println("da vinci")
	}

}
