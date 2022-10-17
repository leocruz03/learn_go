// antes de empezar en go se necesita primero declarar un package
package main

import "fmt"

// funciones en go
// la función main es la que se va a encargar de ejecutar el codigo
func main() {
	// para imprimir en consola se usa el package fmt.Println
	fmt.Println("Hello world in go")

	// declaración de variables, constates y zero values
	const pi float64 = 3.14
	fmt.Println(pi)

	base := 12 // los dos puntos hacen referencia a que esa variable no a sido declarada antriormente
	var altura int = 14
	var area int = 4
	fmt.Println(base, altura, area)

	// el zero value, es un valor que va a tener por defecto una var que no le asignamos un valor
	var a int     // 0
	var b float64 // 0
	var c string  // ""
	var d bool    // false

	fmt.Println(a, b, c, d)

	// operadores aritmeticos
	x := 10
	y := 8

	// suma
	result := x + y
	fmt.Println(result)

	// resta
	result = x - y
	fmt.Println(result)

	// multiplicación
	result = x * y
	fmt.Println(result)

	// división
	result = x / y
	fmt.Println(result)

	// módulo
	result = x % y
	fmt.Println(result)

	// incremental
	x++
	fmt.Println(x)

	// decremental
	x--
	fmt.Println(x)

	// valores primitivos

	// int
	/*
		int8 = 8 bits = -128 a 127
		int16 = 16 bits = -2^15 a 2^15 - 1
		int32 = 32 bits = -2^31 a 2^31 - 1
	*/

}
