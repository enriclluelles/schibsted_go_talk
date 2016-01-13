package main

import (
	"sync"

	. "github.com/enriclluelles/schibsted_go_talk/code/pokemon"
)

//start main OMIT
func main() {
	ch := make(chan *Pokemon)

	var wg sync.WaitGroup

	go func() {
		wg.Wait() // HL
		close(ch) // HL
	}()

	f := func(p *Pokemon) {
		ch <- p
		wg.Done()
	}

	wg.Add(8)
	GetPokemons(f, 1, 79, 25, 4)
	GetPokemons(f, 39, 37, 63, 17)

	for p := range ch {
		p.Print()
	}
}

//end main OMIT
