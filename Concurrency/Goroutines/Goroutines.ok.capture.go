// https://www.youtube.com/watch?v=icbFEmh7Ym0
// https://www.golang-book.com/books/intro/10
// https://tour.golang.org/concurrency/1

package main

import (
	"fmt"
	"sync"
)

func main() {
	data, wg := "hello", sync.WaitGroup{}
	wg.Add(1)
	go exec(data, &wg)
	data = "goodby" // be aware of capture variables
	wg.Wait()
}

func exec(data string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(data)
}
