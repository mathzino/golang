package main

import (
	"fmt"
	"runtime"
	"time"
)

func sendData(ch chan<- int) {
    // randomizer := rand.New(rand.NewSource(time.Now().Unix()))

    for i := 0; true; i++ {
        ch <- i
        time.Sleep(1 * time.Millisecond)
    }
}


func retreiveData(ch <-chan int) {
    loop:
    for {
        select {
        case data := <-ch:
            fmt.Print(`receive data "`, data, `"`, "\n")
        case <-time.After(time.Second * 5):
            fmt.Println("timeout. no activities under 5 seconds")
            break loop
        }
    }
}

func main() {
    runtime.GOMAXPROCS(2)

    var messages = make(chan int)

    go sendData(messages)
    retreiveData(messages)
}