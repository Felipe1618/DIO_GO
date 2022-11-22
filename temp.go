package main

import "fmt"

const ebulicaoF = 212. //fora de um code blocks

func main() {
	tempF := ebulicaoF // Operador curto := -> gopher -> marmota
	// := só é possivel usa-lo dentro de um code blocks
	tempC := (tempF - 32) * 5 / 9 //conversão de F -> C
	fmt.Println("A temperatura de ebulição da água em ºF = ", tempF)
	fmt.Println("A temperatura de ebulição da água em ºC = ", tempC)

	fmt.Printf("A temperatura de fusão da água em ºC: %v (%T) e em ºF: %v (%T)", tempC, tempC, tempF, tempF)
}
