package main

import "fmt"

func main() {

	//for i := 0; i <= 10; i++ {
	//	if i == 0 {
	//		fmt.Println("Zero")
	//	} else if i == 1 {
	//		fmt.Printf("Um")
	//	}
	//}

	for i := 0; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("Um")
		case 2:
			fmt.Println("Dois")
		case 3:
			fmt.Println("Três")
		}
	}

}
