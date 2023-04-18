package main

import "fmt"

func main() {

	x := 4

	if x == 2 || x == 3 {
		fmt.Println("X = 2 ou 3.")
	} else if x > 3 {
		fmt.Println("X maior que 3.")
	}

	if x > 2 && x > 3 {
		fmt.Println("X é maior que 2 e 3.")
	}

}
