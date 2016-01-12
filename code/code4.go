package main

import (
	"fmt"
	"strings"
)

type Contact struct {
	Name     string
	Adresses []string
}

//start main OMIT
func (c Contact) Greeting() {
	a := c.Adresses
	s := strings.Join(a[:len(a)-1], ", ")
	fmt.Printf("Hello there, I'm %s\n", c.Name)
	fmt.Printf("I've lived in %s and %s", s, a[len(a)-1])
}

func main() {
	jedi := Contact{"Luke Skywalker", []string{"tatooine", "dagobah", "endor", "ahch-to"}}
	jedi.Greeting()
}

//end main OMIT
