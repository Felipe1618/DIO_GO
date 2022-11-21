package main

import "fmt"

var a int = 20
var b string = "Nome"

func main() {
	var c float64 = float64(a) //tipo(variavel)
	fmt.Printf("O valor de c é: %g, e o valor de B é: %s.", c, b)
}
