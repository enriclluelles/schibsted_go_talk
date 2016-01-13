package main

import "fmt"

type MyType struct {
	Content string
}

func main() {
	// start main OMIT
	var intCh chan int
	intCh = make(chan int)

	intCh <- 1
	intCh <- 2

	res := <-intCh
	fmt.Println("this is the result: %d", res)
	// end main OMIT

}
