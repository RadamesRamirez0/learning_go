package main

import "fmt"

func main() {
	var opcion int

	for opcion != 5 {

		MostrarMenu()

		opcion = LeerOpcion()

		if opcion == 5 {
			fmt.Println("¡Adios!")
			return
		}

		if opcion < 1 || opcion > 5 {
			fmt.Println("Opción no válida")
			continue
		}

		var valor1 int
		var valor2 int

		valor1, valor2 = LeerValores()

		var calculadora Calculadora
		calculadora.opcion = opcion
		calculadora.valor1 = valor1
		calculadora.valor2 = valor2

		fmt.Println(calculadora.HacerOperacion())

	}

}
