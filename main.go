package main

import "fmt"

const helloWorld string = "Hola %s %s, bienvenido a go \n" //Se declara una constante indicando tipo de dato
const testConst = "Test"                                   //tambien se puede declarar la constante sin indicar el tipo de dato

func main() {
	var name string //Forma recomendada para declarar variables. Se especifica el tipo de dato
	name = "Nombre" //se asigna el valor "Nombre" a la variable name.

	lastName := "<Modificar con el apellido>" //forma corta de declarar variables, sin necesidad de declarar el tipo de dato

	var number = 100 //otra forma de declarar variables sin tipo de dato, pues go automaticamente detecta que tipo de dato es.

	var (
		a = 1 //Forma de declarar multiples variables de manera simultanea.
		b = 2
		c = 3
	)

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)                 //Solicita interaccion del usuario
	fmt.Printf(helloWorld, name, lastName) //llama la constante helloWorld y recibe adicionalmente la variable lastName
	fmt.Println("Hola Mundo")              //igual a fmt.Print pero con salto de linea al final
	fmt.Println(number, a, b, c)
}

func getName() string { //Formato recomendado para declaracion de funciones.  func + nombreFuncion() + Lo que se espera devolver (en este caso, string)

}
