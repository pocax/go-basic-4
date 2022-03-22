package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Main execution started")

	go firstProcess(8)

	secondProcess(8)

	fmt.Println("No. of Go Routines:", runtime.NumGoroutine())
	time.Sleep(time.Second * 2)
	fmt.Println("Main execution ended")
}

func firstProcess(index int) {
	fmt.Println("First process started")
	for i := 1; i < index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("First process ended")
}

func secondProcess(index int) {
	fmt.Println("Second process started")
	for j := 0; j < index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("Second process ended")
}
