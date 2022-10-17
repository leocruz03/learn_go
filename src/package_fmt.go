package main

import "fmt"

func main() {
	helloMessage := "hello"
	worldMessage := "world"

	// println, sirve para imprimir por consola con un salto de linea
	fmt.Println(helloMessage, worldMessage)

	// printf, aparte de imprimir agrega una función extra el %s es para str y %d para números y el %v se utiliza si no se sabe cual podría ser el dato que devuelve la var o const
	nombre := "platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)

	// sprintf, el genera un string no lo imprime, solo lo guarda
	completeMessage := fmt.Sprintf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Println(completeMessage)

	// el tipo de variable con printf %T
	fmt.Printf("message: %T\n", helloMessage)
}
