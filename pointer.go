package main

import (
	"fmt"
)

func main() {
	var a *int
	var b int = 5
	var c int = 6
	fmt.Println("a", a)
	fmt.Println("b", &b)
	fmt.Println("c", &c)
	a = &b
	a = &c
	fmt.Println("a", a)
	fmt.Println("b", &b)
	fmt.Println("c", &c)

	fmt.Println("av", *a)
	fmt.Println("aav", &a)
}
