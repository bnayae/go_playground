// https://www.youtube.com/watch?v=icbFEmh7Ym0
// https://www.golang-book.com/books/intro/10
// https://tour.golang.org/concurrency/1

package main

import (
	"fmt"
	"time"
)

func main() {
	data := "hello"
	go func(d string) {
		fmt.Println(d)
	}(data)
	data = "goodby" // be aware of capture variables
	time.Sleep(100 * time.Microsecond)
}
