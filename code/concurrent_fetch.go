package main

import (
	"time"

	. "github.com/enriclluelles/schibsted_go_talk/code/pokemon"
)

//start main OMIT
func main() {

	f := func(p *Pokemon) {
		p.Print()
	}

	GetPokemons(f, 1, 79, 25, 4)
	GetPokemons(f, 39, 37, 63, 17)

	time.Sleep(10 * time.Second)
}

//end main OMIT
