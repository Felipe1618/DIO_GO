package main

import "fmt"

const ebulicaoF = 212.0

func main() {
	tempF := ebulicaoF
	tempC := (tempF - 32) * 5 / 9
	fmt.Printf("A temperatura de ebulição da água em ºF = %g. A temperatura de ebulição da água em ºC = %g", tempC, tempF)
}
