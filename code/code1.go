package main

import "fmt"

func prependThisIs(s string) string {
	return fmt.Sprintf("this is %s", s)
}

func main() {
	var num int = 3
	str := "sparta"
	res := prependThisIs(str)
	for i := num; i > 0; i-- {
		fmt.Println(res)
	}
}
