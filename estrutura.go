package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func main() {
	fmt.Println(pessoa{"Pedro", 10})
	fmt.Println(pessoa{"Ana", 15})
}
