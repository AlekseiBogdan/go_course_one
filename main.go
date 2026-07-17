package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(square(a))
}

func square(n int) int {
	return n * n
}
