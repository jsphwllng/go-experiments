<div align="center">
<a href="https://github.com/egonelbre/gophers" rel="some text">
	<img style="size: 30%;" src="https://raw.githubusercontent.com/egonelbre/gophers/master/.thumb/animation/gopher-dance-long-3x.gif" alt="welcome to kmy github!">
	</a>
</div>

# A place for me to document my Go learnings
My end goal is to write a program that will take a photograph through a webcam daily and then upload this picture to instagram. I have a long way to "go" before this though.

## resources

* [a tour of go](https://tour.golang.org)
* [go by example](https://gobyexample.com)
* [a cool talk from google about concurrency in go](https://www.youtube.com/watch?v=f6kdp27TYZs)
* [a nice tutorial series from sentdex ](https://www.youtube.com/playlist?list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX) (i would recommend this AFTER the tour of go)
* [a good collection of medium articles](https://medium.com/rungo?source=post_page-----79b82836838b--------------------------------)
* [the go programming language -  Donovan Alan A. A., Kernighan Brian W.](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing-ebook/dp/B0184N7WWS)

## Thoughts on Go

* having to declare the types of parameters and the return type is pretty interesting
* defer is a good concept and could make human-reading functions a bit quicker
* the &variable and *variable pointers are interesting but it seems to boil down to scope which is always annoying
* structures are interesting
* uninitialised variables always being the `0` value (i.e false or 0) is interesting 
* functions behave normally _i.e you can use them to return values_
* pointers: & refers to the address and * to the value stored at the address. For performance reasons you may use pointers as arguments in functions.

* if you specify the return variable name you can iterate on it in the function

```go

func sum(values ...int) (result int) {
	for _, v := range(values)
		result += v
	return	
}
```

* anonymous functions - the last () calls the function

```go
func main() {
	func() {
		fmt.Println("Hello!)
	}() 
}
```

* structs

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
	fmt.Printf("%+v", joe.salary)
```

* comparing structs

```go
type Point struct{ x, y int }

func main() {
	a := Point{1, 2}
	b := Point{1, 2}
	c := Point{2, 1}
	d := Point{y: 2, x: 1}
	fmt.Println("a == b: ", a == b) //true
	fmt.Println("b == c: ", b == c) //false
	fmt.Println("c == d: ", c == d) //false
	fmt.Println("d == a: ", d == a) //true
}
```

* maps are essentially dictionaries in python which is fine just iterating over them is a bit more tricky **maps are also unordered** and black keys return `0` however we can get around this with the `ok` check

```go

testMap := map[string]int{
	"a string key": 1337,
	"paired with": 96,
	"an int value": 11111,
}

for key, value := range testMap {
		fmt.Println(key, "=>", value)
	}
	//a string key => 1337...
fmt.Println(testMap["an int value"])
	// 11111
delete(testMap, "paired with")	
if value, ok := testMap["paired with"]; ok {
	fmt.Println(value)
} else {
	fmt.Println("not a key")
}
//not a key
```

*Converting a Go data stuct JSSON is called marshaling.

```go
data, err := json.Marshal(exampleMap)
if err != nil {
	log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
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

* slices behave interestingly
_slices describe a section of an array so editing a section edits the array_

```go
array := [4]int{1,2,3,4,}
fmt.Println(array)
//[1, 2, 3, 4]
slice := array[0:2]
fmt.Println(slice)
//[1 2]
slice[0] = 1337
fmt.Println(slice, array)
//[1337 2] [1337, 2, 3, 4]
```

* iterating over an array

```go
array := strings.Split("meow", "")
for index, value := range array{
	fmt.Println(index, value) //m, e, o, w
}
```

* i _hate_ how you have to do array[-1] in golang

```go
var a [...]int{1, 2, 3, 4, 5}
fmt.Println(a[len(a)-1])
```

* an interesting way to make an array of runes

```go
var array = []rune("joseph")
	for _, v := range array {
		fmt.Println(string(v)) // j o s e p h
	}
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

* go does **not** support inheritance however you can 'embed' structs within themselves: the bird here has an animal but it is _not_ an animal.

```go
type Animal struct {
	name string
	origin string
}

type Bird struct {
	Animal
	speed	float32
	canFly	bool
}

func main(){
	b := Bird{}
	b.name = "Dave"
	b.origin = "NZ"
	b.speed = 90
	b.canFly = true
	fmt.Printf("%+v\n", b)
	fmt.Println(b.name)
	//{Animal:{name:Dave origin:NZ} speed:90 canFly:true}
	//Dave
	bb := Bird{
		Animal: Animal{name: "mike", origin: "Auz"},
		speed: 200,
		canFly: true,
	}
	fmt.Println(bb.name)
	//mike
	fmt.Println(bb.Animal)
	//{mike Auz}
}
```

* defer statements:

```go
fmt.Println("1")
defer fmt.Println("2")
fmt.Println("3")
//1
//3
//2
```

* iota is a 'Constant Generator' beginning at zero 

```go
type Dog int

const (
	Layla Dog = iota
	Winston
	Shiro
	Angel
	Snoop
	Priscilla
	Nate
)

func main() {
	fmt.Println(Shiro) //2
	fmt.Println(Layla) //0
}
```

# Examples
### fizzbuzz
```go
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("fizz")
		} else if i%3 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("done!")
}
```
### Using multiple, unnamed variables in a function

```go
func main() {
	scream("hello", "goodbye", "nihao", "salam", "oh", "aaaaii")
}


func scream(words ...string) {
	for _, v := range words {
		fmt.Println(strings.ToUpper(v) + "!!!!!!!!!!!!!")
	}
}
```
