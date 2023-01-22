package main

import (
	"fmt"
	"time"
)

func main() {
	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("main function uses concurrency with f function")
}

func f() {
	fmt.Println("f function")
}
