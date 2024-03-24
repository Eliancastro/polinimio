
package main

import "fmt"


func mostrarPolinomio(coeficientes []float64) {
	grado := len(coeficientes) - 1

	fmt.Printf("%.1f", coeficientes[0])

	for i := 1; i < len(coeficientes); i++ {
		fmt.Printf(" + %.1f x^%d", coeficientes[i], i)
	}

	fmt.Printf("\n")
}

func main() {
	coeficientes := []float64{3.0, 2.0, 1.0}
	mostrarPolinomio(coeficientes)
}


/**En este código, la función
mostrarPolinomio
recibe un slice de coeficientes de tipo
float
y muestra por pantalla el polinomio completo. Luego, en la función
main, se define un slice de coeficientes
{3.0, 2.0, 1.0}
y se llama a la función mostrarPolinomio
con esos coeficientes. Esto imprimirá en pantalla el polinomio completo
3.0 + 2.0 x + 1.0 x^2**/