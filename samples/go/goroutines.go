package main

import (
	"fmt"
	"sync"
)

func f(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go f(i, &wg)
	}
	wg.Wait()
}
