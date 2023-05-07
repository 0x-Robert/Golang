package main

import (
	"fmt"
	"sort"
)

func main() {
    var n int
    fmt.Scanln(&n)
    a := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&a[i])
    }
    p := make([]int, n)
    for i := 0; i < n; i++ {
        p[i] = i
    }
    sort.Slice(p, func(i, j int) bool { return a[p[i]] < a[p[j]] })
    b := make([]int, n)
    for i := 0; i < n; i++ {
        b[p[i]] = i
    }
    for i := 0; i < n; i++ {
        fmt.Print(b[i], " ")
    }
}