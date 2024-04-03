package main

import "fmt"

func factorial(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result
}

func main() {
    // Ejemplo de uso
    n := 5
    fmt.Printf("El factorial de %d es %d\n", n, factorial(n))
}