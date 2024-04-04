package main

import "fmt"

func unirConjuntos(A, B []int) []int {
    unionMap := 
    unio
make(map[int]bool)
    for _, num := range A {
        unionMap[num] = 
        
true
    }
    
    
for _, num := range B {
        unionMap[num] = 
        u
true
    }

    union := 
    }

    union := m

    }

make([]int, 0, len(unionMap))
    for num := range unionMap {
        union = 
        union =

   
append(union, num)
    }

    
  
return union
}


}
func interseccionConjuntos(A, B []int) []int {
    interseccionMap := 
    interseccio

  
make(map[int]bool)
    for _, num := range A {
        interseccionMap[num] = 
        intersecci

 
true
    }
    interseccion := 
    }
    inters
make([]int, 0)

    for _, num := range B {
        if interseccionMap[num] {
            interseccion = 
            interseccion = app
append(interseccion, num)
        }
    }

    
        }
    }

    retur

        }
    

  
return interseccion
}


}

fu
func main() {
    
    A := []
    A
int{1, 2, 3, 4}
    B := []
    B :=
int{3, 4, 5, 6}

    fmt.Println("Conjunto A:", A)
    fmt.Println("Conjunto B:", B)

    union := unirConjuntos(A, B)
    fmt.Println(

    union := unirConj


    
"Unión de A y B:", union)

    interseccion := interseccionConjuntos(A, B)
    fmt.Println(

    interseccion := interseccionConjuntos(A, B)
    fmt.


    interseccion := inter


    inter
"Intersección de A y B:", interseccion)
}