package main

import (
	. "github.com/enriclluelles/schibsted_go_talk/code/pokemon"
)

//start main OMIT
func main() {
	ch := make(chan *Pokemon)

	f := func(p *Pokemon) {
		ch <- p
	}

	GetPokemons(f, 1, 79, 25, 4)
	GetPokemons(f, 39, 37, 63, 17)

	for p := range ch {
		p.Print()
	}
}

//end main OMIT
