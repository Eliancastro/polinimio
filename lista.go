package main

import (
"fmt"
"math"
)

func encontrarMinMax(lista []int) (int, int) {
min := math.MaxInt32
max := math.MinInt32

for _, num := range lista {
if num < min {
min = num
}
if num > max {
max = num
}
}

return min, max
}

func main() {
numeros := []int{10, 5, 8, 3, 12, 6}
minimo, maximo := encontrarMinMax(numeros)
fmt.Printf("El menor número es: %d\n", minimo)
fmt.Printf("El mayor número es: %d\n", maximo)
}