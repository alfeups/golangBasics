package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Número par: %d \n", i)
		} else {
			fmt.Printf("Número impar: %d \n", i)
		}
	}

}
