package main

import (
	"fmt"
)

func main() {
	var opcion int

	fmt.Println("Menú:")
	fmt.Println("1. Opción 1")
	fmt.Println("2. Opción 2")
	fmt.Println("3. Opción 3")
	fmt.Println("4. Opción 4")

	fmt.Print("Elige una opción: ")
	fmt.Scanln(&opcion)

	switch opcion {
	case 1:
		fmt.Println("Has elegido la Opción 1")
	case 2:
		fmt.Println("Has elegido la Opción 2")
	case 3:
		fmt.Println("Has elegido la Opción 3")
	case 4:
		fmt.Println("Has elegido la Opción 4")
	default:
		fmt.Println("Opción incorrecta. Por favor, elige una opción válida.")
	}
}