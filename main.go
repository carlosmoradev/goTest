package main

import "fmt"

const helloWorld string = "Hola %s %s, bienvenido a go \n" //Se declara una constante indicando tipo de dato
const testConst = "Test"                                   //tambien se puede declarar la constante sin indicar el tipo de dato

func main() {
	lastName := "<Modificar con el apellido>" //forma corta de declarar variables, sin necesidad de declarar el tipo de dato
	name := getName()                         //Se declara la variable name llamandola desde la funcion getName()
	var number = 100                          //otra forma de declarar variables sin tipo de dato, pues go automaticamente detecta que tipo de dato es.

	// var (
	// 	a = 1 //Forma de declarar multiples variables de manera simultanea.
	// 	b = 2
	// 	c = 3
	// )

	a, b, c := getVariables()

	fmt.Printf(helloWorld, name, lastName) //llama la constante helloWorld y recibe adicionalmente la variable lastName
	fmt.Println("Hola Mundo")              //igual a fmt.Print pero con salto de linea al final
	fmt.Println(number, a, b, c)
}

func getName() string { //Formato recomendado para declaracion de funciones.  func + nombreFuncion() + Lo que se espera devolver (en este caso, string)
	var name string
	name = "Nombre" //se asigna el valor "Nombre" a la variable name para poder reutilizarlo por medio de la funcion getName().
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name) //Solicita interaccion del usuario
	return name
}

func getVariables() (int, int, int) {
	return 1, 2, 3
}
