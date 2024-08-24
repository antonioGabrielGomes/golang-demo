package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("worker %d got %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {

	/*
		ex 2
	*/
	ch := make(chan int)
	qtdWorkers := 15

	for i := range qtdWorkers {
		go worker(i, ch)
	}

	for i := range 15 {
		ch <- i
	}

	/*
		ex 1
		// goroutine 1
		ch := make(chan string) // empty

		// goroutine 2
		go func() {
			ch <- "Full Cycle"
		}()

		msg := <-ch
		fmt.Println(msg)
	*/
}
