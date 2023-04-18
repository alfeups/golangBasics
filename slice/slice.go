package main

import "fmt"

func main() {

	arr := [7]float64{1, 2, 3, 4, 5, 6, 7}
	fatia := arr[2:5]
	fmt.Println(arr)
	fmt.Println(fatia)

	fatia1 := []int{1, 2, 3}
	fatia2 := append(fatia1, 4, 5)
	fmt.Printf("Fatia 1: %d e Fatia 2: %d", fatia1, fatia2)

}
