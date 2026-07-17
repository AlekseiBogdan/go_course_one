package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a%15 == 0:
		fmt.Println("FizzBuzz")
	case a%5 == 0:
		fmt.Println("Buzz")
	case a%3 == 0:
		fmt.Println("Fizz")
    default:
        fmt.Println(a)
	}
}
