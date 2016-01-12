package main

import (
	"fmt"
	"strings"
)

//start def OMIT
type Quacker interface {
	Quack()
}

//end def OMIT

//start main OMIT
type Duck struct {
}

func (d *Duck) Quack() {
	fmt.Println("quack, I'm a duck")
}

type Geese struct {
}

func (g *Geese) Quack() {
	fmt.Println("quack, I'm a geese")
}

func MakeItQuack(q Quacker) {
	q.Quack()
}

//end main OMIT

//start a OMIT
type MyType struct {
	MyValue string
}

func (mt *MyType) RemoveA() {
	mt.MyValue = strings.Replace(mt.MyValue, "a", "", -1)
}

var i interface{} = MyValue{}

//end a OMIT

//start b OMIT
//won't compile
func LengthWithoutA(mt interface{}) int {
	mt.RemoveA()
	return len(mt)
}

func main() {
	LengthWithoutA(i)
}

//end b OMIT

//start c OMIT
//will compile
func LengthWithoutA(mt interface{}) int {
	mt.(MyType)
	mt.RemoveA()
	return len(mt)
}

func main() {
	LengthWithoutA(i)
}

//end c OMIT
