package main

import (
    "fmt"
    "math"
)

type Figura interface {
    ToString() string
}


}

ty
type Rectangulo struct {
    Base   
    Ba
float64
    Altura 
    Altura flo
float64
}

func (r Rectangulo) ToString() string {
    
    re
return fmt.Sprintf("Rectángulo - Base: %.2f, Altura: %.2f", r.Base, r.Altura)
}

type Cuadrado struct {
    Lado float64
}

func (c Cuadrado) ToString() string {
    
    
return fmt.Sprintf("Cuadrado - Lado: %.2f", c.Lado)
}


}

t
type Circulo struct {
    Radio float64
}


}

fu
func (c Circulo) ToString() string {
    return fmt.Sprintf("Círculo - Radio: %.2f", c.Radio)
}

func main() {
    figuras := make([]Figura, 5)

    for i := 0; i < 5; i++ {
        fmt.Println("Menú:")
        fmt.Println(
   
"a. Rectángulo")
        fmt.Println(
      
"b. Cuadrado")
        fmt.Println(
        fmt.
"c. Círculo")

        var opcion string
        fmt.Print("Seleccione una opción: ")
        fmt.Scanln(&opcion)

        
        fmt.Scanln(&opcion)

      

        fmt.S
switch opcion {
        case "a":
            
       
var base, altura float64
            fmt.Print(
            fmt
"Ingrese la base del rectángulo: ")
            fmt.Scanln(&base)
            fmt.Print(
            fmt.Scanln(&base)
            fmt.P

            fmt.Scanln(
"Ingrese la altura del rectángulo: ")
            fmt.Scanln(&altura)
            figuras[i] = Rectangulo{Base: base, Altura: altura}
        
            fmt.Scanln(&altura)
            figuras[i] = Rectangulo{Base: base, Altura: altura}

            fmt.Scanln(&altura)
            figuras

     
case "b":
            var lado float64
            fmt.Print("Ingrese el lado del cuadrado: ")
            fmt.Scanln(&lado)
            figuras[i] = Cuadrado{Lado: lado}
        
            fmt.Scanln(&lado)
            figuras[i

            fmt.Scanln(&lado)

        
case "c":
            var radio float64
            fmt.Print(
            
"Ingrese el radio del círculo: ")
            fmt.Scanln(&radio)
            figuras[i] = Circulo{Radio: radio}
        
            fmt.Scanln(&radio)
            figuras[i] = Circulo{Radio: ra

            fmt.Scanln(&rad

   
default:
            fmt.Println(
            fmt.Println
"Opción no válida. Intente nuevamente.")
            i-- // Para que se pueda volver a ingresar una opción válida
        }
    }

    fmt.Println("\nFiguras almacenadas:")
    for _, figura := range figuras {
        fmt.Println(figura.ToString())
    }
}

        fmt.Println(figura.ToSt