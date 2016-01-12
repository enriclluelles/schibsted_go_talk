package main

import (
	"fmt"
	"strings"
)

//start OMIT
var num int = 1 //globals FTW

func main() {
	str := "this is sparta"

	split := func(s string) (string, string) {
		var splitted []string = strings.Split(s, " ")
		return splitted[0], splitted[1] // multiple returns! // HL
	}

	var first, second string = split(str)
	fmt.Println("what")
	fmt.Println(second)
	fmt.Println(first)
	fmt.Println("I don't even")
}

//end OMIT
