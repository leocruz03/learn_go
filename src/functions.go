// uso de funciones anonimas y comunes

package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func main() {
	// invocar funciones
	normalFunction("Hello")
	tripleArgument(1, 2, "Leonardo")

	value := returnValue(4)
	fmt.Println(value)
}
