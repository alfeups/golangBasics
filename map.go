package main

import "fmt"

func main() {

	x := make(map[string]int)
	x["Key"] = 10 //key - value
	fmt.Println(x)
	fmt.Println(x["Key"])

	y := make(map[int]int)
	y[1] = 10
	y[2] = 20
	fmt.Println(y[1])
	fmt.Println(y[2])
	fmt.Println("Y1 + Y2 =", y[1]+y[2])
}
