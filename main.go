package main

import (
	"fmt"
	"strconv"
    "strings"
)

func main() {
	var a string
	fmt.Scan(&a)
    line := strings.TrimSpace(a)
	x, err := strToInt(line)
    if err != nil {
        fmt.Println("bad")
    } else {
        fmt.Printf("ok %d", x)
    }

}

func strToInt(line string) (int, error) {
    n, err := strconv.Atoi(line)
	return n, err
}
