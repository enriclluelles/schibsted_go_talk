package main

import "fmt"

//start contact OMIT
type Contact struct {
	Name     string
	Adresses []string
}

//end contact OMIT

//start main OMIT
func modify(c Contact) { // HL
	c.Name = "Darth Sidious"
	c.Adresses = []string{"naboo"}
}

func modify2(c *Contact) { // HL
	// c is a pointer but no need to dereference
	// to access its fields
	c.Name = "Qui-Gon Jinn"
	c.Adresses = []string{"coruscant"}
}

func main() {
	jedi := Contact{"Luke Skywalker", []string{"tatooine", "dagobah"}}
	modify(jedi) // HL
	fmt.Println("%#v", jedi)
	modify2(&jedi) // HL
	fmt.Println("%#v", jedi)
}

//end main OMIT
