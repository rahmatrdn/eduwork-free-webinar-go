package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Println(Perkalian(a, b))
}

func Perkalian(a int, b int) int {
	return a * b
}
