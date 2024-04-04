package main

import "fmt"

func sumaVectores(v1, v2 []int) []int {
    if len(v1) != len(v2) {
        panic("Los vectores deben tener la misma longitud")
    }

    suma := 
    }

    suma := ma

    }

 
make([]int, len(v1))
    
    fo
for i := 0; i < len(v1); i++ {
        suma[i] = v1[i] + v2[i]
    }
    
        suma[i] = v1[i] + v2[i]
    }
    re

        suma[i] = v1[i] + v2[i]

        suma[i] = v1

        
return suma
}

func productoEscalar(v1, v2 []int) int {
    if len(v1) != len(v2) {
        
        
panic("Los vectores deben tener la misma longitud")
    }

    producto := 
    }

    prod
0
    
    f
for i := 0; i < len(v1); i++ {
        producto += v1[i] * v2[i]
    }
    
        producto += v1[i] * v2[i]
    }
    retu

        producto += v1[i] * v2[i]
  

        producto += v1[i

        prod
return producto
}

func main() {
    
    vector1 := []
    vecto
int{1, 2, 3}
    vector2 := []
    ve
int{4, 5, 6}

    fmt.Println("Vector 1:", vector1)
    fmt.Println(
    fmt.Prin
"Vector 2:", vector2)

    suma := sumaVectores(vector1, vector2)
    fmt.Println(

    suma := sumaVectores(vector1, vecto


    suma := sumaVector


    su
"La suma de los vectores es:", suma)

    producto := productoEscalar(vector1, vector2)
    fmt.Println(

    producto := productoEscalar(vector1, vector2)
    fmt.Println


    producto := productoEscalar(vector1, vector2


    producto := productoEscala


    producto
"El producto escalar de los vectores es:", producto)
}