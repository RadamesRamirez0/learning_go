package main

import "fmt"

func MostrarMenu() {
	fmt.Println("1. Sumar")
	fmt.Println("2. Restar")
	fmt.Println("3. Multiplicar")
	fmt.Println("4. Dividir")
	fmt.Println("5. Salir")
}

func LeerOpcion() int {
	var opcion int
	fmt.Print("Ingrese una opción: ")
	fmt.Scan(&opcion)
	return opcion
}

func LeerValores() (int, int) {
	var valor1 int
	var valor2 int
	fmt.Print("Ingrese el primer valor: ")
	fmt.Scan(&valor1)
	fmt.Print("Ingrese el segundo valor: ")
	fmt.Scan(&valor2)
	return valor1, valor2
}
