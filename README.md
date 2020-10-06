# A place for me to document my Go learnings
<p>My end goal is to write a program that will take a photograph through a webcam daily and then upload this picture to instagram. I have a long way to "go" before this though.</p>

<div align="center">
<a href="https://github.com/egonelbre/gophers" rel="some text">
	<img style="size: 30%;" src="https://raw.githubusercontent.com/egonelbre/gophers/master/.thumb/animation/gopher-dance-long-3x.gif" alt="welcome to my github!">
	</a>
</div>

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

* go appends arrays by making new arrays for example

```go
a := make([]string, 2)
s := append(a, "test")
fmt.Println(a)
// []
fmt.Println(s)
// [ test]
```