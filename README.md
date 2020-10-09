<div align="center">
<a href="https://github.com/egonelbre/gophers" rel="some text">
	<img style="size: 30%;" src="https://raw.githubusercontent.com/egonelbre/gophers/master/.thumb/animation/gopher-dance-long-3x.gif" alt="welcome to kmy github!">
	</a>
</div>

# A place for me to document my Go learnings
<p>My end goal is to write a program that will take a photograph through a webcam daily and then upload this picture to instagram. I have a long way to "go" before this though.</p>

## resources

* [a tour of go](https://tour.golang.org)
* [go by example](https://gobyexample.com)
* [a cool talk from google about concurrency in go](https://www.youtube.com/watch?v=f6kdp27TYZs)
* [a nice tutorial series from sentdex ](https://www.youtube.com/playlist?list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX) (i would recommend this AFTER the tour of go)
* [a good collection of medium articles](https://medium.com/rungo?source=post_page-----79b82836838b--------------------------------)

## Thoughts on Go

* having to declare the types of parameters and the return type is pretty interesting
* defer is a good concept and could make human-reading functions a bit quicker
* the &variable and *variable pointers are interesting but it seems to boil down to scope which is always annoying
* structures are interesting

```go
type Salary struct {
	basic, allowance, bonuses int
}

type Employee struct {
	firstName, lastName string
	salary              Salary
	fullTime            bool
}


joe := Employee{
		firstName: "joe",
		lastName:  "phwllng",
		salary:    Salary{10, 20, 30},
		fullTime:  false,
	}
	fmt.Println(joe.salary.basic)
	joe.salary.basic = 20
	fmt.Println(joe.salary.basic)

```
* maps are essentially dictionaries in python which is fine just iterating over them is a bit more tricky

```go
for key, value := range testMap {
		fmt.Println(key, "=>", value)
	}
```

* go appends arrays by making new arrays for example:

```go
a := make([]string, 2)
s := append(a, "test")
fmt.Println(a)
// []
fmt.Println(s)
// [ test]
```

* syncs and wait groups

```go

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go say("hello!")
	wg.Add(1)
	go say("neato")
	wg.Wait()
}


```
* channels

```go

// channel <- data
//         ^ pushes data to a channel
// <- channel
// ^ we read the channel for example

c := make(chan string)
c <- "hello"
channelString := <- c
fmt.Println(channelString)
//"hello"


```

* building a _simple_ web server seems easy too!

```go
package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//fprint will format will specify based on the writer and will output hello
	fmt.Fprintf(w, "<h1 style='font-size: 200px'>hello!</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	//fprint will format will specify based on the writer and will output hello
	fmt.Fprintf(w, "i am joe!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.ListenAndServe(":3000", nil)
}
```
* parsing a website isn't hard either

```go
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//fprint will format will specify based on the writer and will output hello
	fmt.Fprintf(w, "<h1 style='font-size: 200px'>hello!</h1>")
	resp, _ := http.Get("https://www.lewagon.com/")
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)
	fmt.Fprintf(w, stringBody)
}
```