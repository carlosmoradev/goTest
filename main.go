package main

import "fmt"

func main() {
	var name string
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name) //Solicita interaccion del usuario
	fmt.Printf("Hola %s, bienvenido a go \n", name)
	fmt.Println("Hola Mundo") //igual a fmt.Print pero con salto de linea al final
}
