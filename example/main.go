package main

import (
	"fmt"
	"time"

	"github.com/guil95/semaphore"
)

func main() {
	quit := make(chan error)

	go semaphore.Run(10, doSomething, quit)

	for {
		select {
		case q := <-quit:
			panic(q)
		}
	}
}

func doSomething() error {
	fmt.Println("started process")
	time.Sleep(time.Second * 3)
	fmt.Println("finished process")

	return nil
}
