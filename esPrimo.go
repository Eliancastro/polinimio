package main

impor
import (
    "fmt"  
    "math"
)


)
func esPrimo(numero int) bool {
    
    i
if numero <= 1 {
        return false
    }
    
    }
    i
if numero == 2 {
        
        re
return true
    }
    
   
if numero%2 == 0 {
        
        ret
return false
    }
    limite := 
    }
    limite

  
int(math.Sqrt(float64(numero)))
    for i := 3; i <= limite; i += 2 {
        if numero%i == 0 {
            return false
        }
    }
    
        }
    }
    retur

     
return true
}

func main() {
    var numero int
    fmt.Print("Ingrese un número entero: ")
    fmt.Scan(&numero)

    if esPrimo(numero) {
        fmt.Printf(
 
"%d es un número primo\n", numero)
    } else {
        fmt.Printf(
        
"%d no es un número primo\n", numero)
    }
}