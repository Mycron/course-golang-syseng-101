# Unit 1: A tour of Go
Please finish all materials on https://go.dev/tour/ or https://go.dev/tour/list except "Generics" and "Concurrency"

Note: if you stuck with module try to read https://go.dev/blog/using-go-modules.

## FAQ

TBA

## Quiz

#### Q1. How will you add the number 3 to the right side?

`numbers := []int{1, 1, 2}`

1. `numbers.append(3)`
2. `numbers.insert(3, 3)`
3. `append(numbers, 3)`
4. `numbers = append(numbers, 3)` - this doing it

#### Q2. From where is the variable fooVar accessible if it is declared outside of any functions in a file in package fooPackage located inside module fooModule

1. anywhere inside `fooPackage`, not the rest of `fooModule` - here
2. by any application that imports `fooModule`
3. from anywhere in `fooModule`
4. by other packages in `fooModule` as long as they import `fooPackage` - and here

#### Q3. What should the idiomatic name be for an interface with a single method and the signature Serve() error

1. Servable
2. Server
3. ServeInterface
4. IServe

#### Q4. Which is **not** valid loop construct?

1. `for i,r:=0,rand.Int(); i < r%10; i++ { ... }`
2. `for { ... }`
3. `{ ... } for false` -- here is a failure
4. `for _,c := range "foobar" { ... }`

## Excercises
`E0` is for illustration how to work and submit Excercises.
Let's make 

### E0. Exercise: Loops and Functions
More: https://go.dev/tour/flowcontrol/8
let's create implementation in `unit1/exercises/e0/main.go`, wherer `unit1` is number of current unit, `e0` is number of current exercise.

If task has multiple steps to do, we assume that the last steps is final. Input values must be the same as in Exercie definition if other is not mentioned. 

At first we've been asked to implement Sqrt function with partial implementation of Newton method and send output of computation steps on each iteration of method:

```go
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i:=0; i<10; i++ {
		delta := (z*z - x) / (2*z)
		z -= delta
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
```

Next we've asked to change loop condition to stop once the value has stopped changing (or only changes by a very small amount). We can achieve that by adding `epsilon` constant that will be compared with `delta`. 
**Note that their absolute values should be compared**. One approach is to use `Abs()` function form `math` package, or we can notice that `x` argument must be positive and care(?) our own condition based on partial comparison of positive-only numbers:

```go
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	epsilon := 0.0000000000001
	z := 1.0
	for {
		delta := (z*z - x) / (2 * z)
        // or: if math.Abs(delta) < epsilon {
		if delta < epsilon && delta > -epsilon {
			break
		}
		z -= delta
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
```

If you haven't managed to solve all steps publish code for the step you succeeded and make a comment in code.
If you exercise requires to make multiple files or even packages, don't hesitate to create them in `unit1/exercises/eX/` folder

As soon as you implemented the code and place it into `unit1/exercises/e0/main.go`, make PR to this repo.

---

### E1. Fibonacci closure
More: https://go.dev/tour/moretypes/26

**Note**: Edit only `fibonacci` function definition:
```go
func fibonacci() func() int {}
```
Keep `main` function untouch. Don't add additional Prints to output. It is checked in tests.

Share your implementation `unit1/exercises/e1/main.go` in github PR.

---

### E2. Stringers
More: https://go.dev/tour/methods/18

Implement `func (ip IPAddr) String() string` method for `IPAddr` structure. 
Note that receiver variable `ip` is not pointer.

Keep `main` function untouch. Don't add additional Prints to output. It is checked in tests.

Share your implementation  `unit1/exercises/e2/main.go` in github PR.

---

### E3. Errors
More: https://go.dev/tour/methods/20

**Note**: errors are just values and must be treated as values. In this scenario error value will be `ErrNegativeSqrt` struct that satisfy error type. Nil is valid error, pointer to struct with `Error() string` method is valid error too.

Keep `main` function untouch. Don't add additional Prints to output. It is checked in tests.

Share your implementation  `unit1/exercises/e3/main.go` in github PR.

---

### E4. rot13Reader
More: https://go.dev/tour/methods/23

`'A'` is a byte. It is treated as 8 bit number, so it can be compared and added/substracted. 

not that A-Z and a-z letters have their own ranges and should be mapped to letter from their group: letter form A-Z must be mapped to letter from A-Z. Same for a-z.

https://www.rapidtables.com/code/text/ascii-table.html


Keep `main` function untouch. Don't add additional Prints to output. It is checked in tests.

Share your implementation  `unit1/exercises/e4/main.go` in github PR.

---

### E5. Images
More: https://go.dev/tour/methods/25

**Note**: read the docs for [golang.org/x/tour/pic](https://pkg.go.dev/golang.org/x/tour/pic) package.

Keep `pic.ShowImage(m)` in `main` function untouch. Don't add additional Prints to output. It is checked in tests. 

Share your implementation  `unit1/exercises/e5/main.go` in github PR.

---