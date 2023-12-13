package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
    for i := 0; i < till; i++ {
        fmt.Println((i + 1), message)
    }
}

func main() {
    runtime.GOMAXPROCS(1)

    go print(5000000, "halo")
    print(5, "apa kabar")

    var input string
    fmt.Scanln(&input)
}