package main

import "fmt"

var ebulicaoKelvin = 373

func main() {
	celsius := ebulicaoKelvin - 273
	fmt.Printf("Temperatura de ebulição da água em Kelvin = %dK e"+
		" em Celsius = %dºC", ebulicaoKelvin, celsius)
}
