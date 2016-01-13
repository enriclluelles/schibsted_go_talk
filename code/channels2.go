package main

type MyType struct {
	Content string
}

func main() {
	// start main OMIT
	intCh := make(chan int)
	stringCh := make(chan string)
	mCh := make(chan *MyType)

	go func() {
		intCh <- 1
	}()

	go func() {
		stringCh <- "from channel!"
	}()

	<-intCh
	<-stringCh
	<-mCh
	// end main OMIT

}
