package main

// import packages together.
import (
	"fmt"
	"math"
)

func multiple(a int, b int) int {
	return a * b
}

func triple(a, b, c int) int {
	return a * b * c
}

// a function that returns multiple result
func swap(a, b int) (int, int) {
	return b, a
}

// a function that returns a named result
func piyo() (piyo string) {
	return "piyo"
}

func main() {
	fmt.Println("Hello, World!")
	// functions provided by a package must be upper camel case.
	fmt.Println(math.Pi)

	fmt.Println(multiple(2, 2))

	// variables
	var (
		a = 5
		b = 2
	)
	fmt.Println(multiple(a, b))

	// variables can be re-assigned.
	a = 10
	fmt.Println(multiple(a, b))

	// constants
	const (
		c = 4
		d = 3
	)
	fmt.Println(triple(a, c, d))

	fmt.Println(func(a, b int) int {
		return a - b
	}(swap(d, c)))

	fmt.Println(piyo())

	// primitives are initialized by default values.
	var (
		num  int
		str  string
		isOk bool
	)
	fmt.Println(num, str, isOk)

	// shorthand declaration.
	num1, num2 := 1, 2
	fmt.Println(num1, num2)

	// for loop and if statement
	sum := 0
	for i := 1; i < 4; i++ {
		sum += i
		if i != 3 {
			fmt.Print(i)

			fmt.Print(", ")
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println(sum)
}
