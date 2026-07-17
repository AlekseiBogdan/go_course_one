package main

import (
	"fmt"
)

func main() {
	var (
		a     int
		total int
	)
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		total += i
	}
	fmt.Println(total)
}
