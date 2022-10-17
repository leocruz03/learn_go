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
		int64 = 64 bits = -2^63 a 2^63 - 1
	*/

	// enteros positivos, cuando sabemos que el datos siempre va a ser positivo se utiliza uint

	/*
		uint8 = 8 bits = 0 a 2^8 - 1
		uint16 = 16 bits = 0 a 2^16 - 1
		uint32 = 32 bits = 0 a 2^32 - 1
		uint64 = 64 bits = 0 a 2^64 - 1
	*/

	// números decimales se utiliza el float

	/*
		float32 = 32 bits = +/- 1.18e^-38 a +/- 3.4e^38
		float64 = 64 bits = +/- 2.23e^-308 a +/- 1.8e^308
	*/

	// string para textos y bool para true o false

	/*
		string = ""
		bool = false/true
	*/

	// números 	complejos, se usa complex

	/*
		complex64 = real o imaginario float32
		complex128 = real o imaginario float64

		ejemplo: c := 10 + 8i
	*/

}
