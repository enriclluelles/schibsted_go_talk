package main

import "fmt"

type Car struct {
}

//start main OMIT
func main() {
	m := make(map[string]*Car)
	if c, ok := m["my car"]; ok {
		fmt.Println("%#v", c)
	}
}

//end main OMIT
