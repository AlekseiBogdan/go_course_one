package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    r := bufio.NewReader(os.Stdin)
    line, _ := r.ReadString('\n')
    parts := strings.Fields(strings.TrimSpace(line))
    nums := make([]int, 0, len(parts))
    max := 0
    for _, p := range parts {
        n, _ := strconv.Atoi(p)
        nums = append(nums, n)
    }
    for _, i := range nums {
        if i > max {
            max = i
        }
    }
    fmt.Println(max)
    _ = fmt.Print
}
