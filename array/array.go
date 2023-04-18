package main

func main() {

	var arr [5]int64

	for i := 0; i < 5; i++ {
		arr[i] = int64(i * 25)
		println(arr[i])
	}
}
