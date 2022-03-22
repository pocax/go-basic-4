package main

import (
	"fmt"
	"sync"
)

func main() {
	fruits := []string{"apple", "banana", "orange"}

	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1)
		go printFruits(index, fruit, &wg)
	}

	wg.Wait()
}

func printFruits(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("Index => %d, Fruit => %s\n", index, fruit)
	wg.Done()
}
