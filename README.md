# NOTAS DE MI APRENDIZAJE DE GO

Este repositorio contiene las notas que he tenido en cuenta para el aprendizaje de este lenguaje de programación.

## Primeros pasos
### Declaración de Constantes
```
const pi float64 = 3.14
const pi2 = 3.16
```
### Declaración de Variables
```
base := 12          //Primera forma
var altura int = 14 //Segunda forma
var area int        //Go no compila si las variables no son usadas
```
### Zero values
Go asigna valores a variables vacías
```
var a int
var b float64
var c string
var d bool
```

## Tipos de Datos Primitivos
### Números Enteros
```
int = Depende del OS (32 o 64 bits)
int8 = 8bits = -128 a 127
int16 = 16bits = -2^15 a 2^15-1
int32 = 32bits = -2^31 a 2^31-1
int64 = 64bits = -2^63 a 2^63-1
```

### Optimizar memoria cuando sabemos que el dato siempre va a ser positivo
```
uint = Depende del OS (32 o 64 bits)
uint8 = 8bits = 0 a 255
uint16 = 16bits = 0 a 2^15-1
uint32 = 32bits = 0 a 2^31-1
uint64 = 64bits = 0 a 2^63-1
```
### Números Decimales
```
float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308
```
### Textos y booleanos
```
string = ""
bool = true or false
```
### Numeros complejos
```
Complex64 = Real e Imaginario float32
Complex128 = Real e Imaginario float64
Ejemplo : c:=10 + 8i
```

## Paquetes principales en GO
### Funciones principales del paquete FMT
```
package main

import "fmt"

func main() {
    helloMessage := "Hello"
    worldMessage := "World"

    // Println Example
    fmt.Println(helloMessage, worldMessage)

    // Printf Example
    nombre := "Nerfo"
    edad := 500
    fmt.Printf("%s tiene %d años", nombre, edad)
    fmt.Printf("%v tiene %v años", nombre, edad)

    // Sprintf Example   
    message := fmt.Sprintf("%s tiene %d años", nombre, edad)
    fmt.Println(message)

    // Tipo de dato
    fmt.Printf("helloMessage: %T", helloMessage)
}
```
## Declaración de Funciones
### Funciones sin parámetros
```
package main

import "fmt"

func primerFuncion() {
    fmt.Println("Hola Mundo")
}

func main() {
    primerFuncion()
}
```
### Funciones con parámetros
```
package main

import "fmt"

func holaMundoFunc(nombre string, edad int) {
    fmt.Sprintf("%s tiene %d años", nombre, edad)
}

func main() {
    holaMundoFunc("Nerfo", 500)
}
```
### Funciones para realizar operaciones
```
package main

import "fmt"

func producto(a int, b int) {
    return a * b
}

func main() {
    value := producto(5, 8)
    fmt.Println("Value: ", value)
}
```
### Conocimiento aplicado
```
package main

import (
	"fmt"
	"math"
)

func areaCirculo(radio float64) float64{ 
	return math.Pi*radio*radio
}
func areaRectangulo(base float64, altura float64) float64 {
	return base*altura
}

func areaTrapezoide(B float64,b float64,h float64) float64{
	return h*(B+b)/2
}

func main() {
	fmt.Printf("Circulo %.2f \n",areaCirculo(2))
	fmt.Printf("Rectangulo %.2f \n",areaRectangulo(5,10))
	fmt.Printf("Trapezoide %.2f \n",areaTrapezoide(10,5,3))
}
```
### Arrays y Slices
```
package main

import "fmt"
func main() {
    // Array
    var array [4]int
    array[0] = 1
    array[1] = 2
    fmt.Println(array, len(array), cap(array))

    // Slice
    slice := []int{0,1,2,3,4,5,6}
    fmt.Println(slice, len(slice), cap(slice))

    // Métodos en el slice
    fmt.Println(slice[0])
    fmt.Println(slice[:3])
    fmt.Println(slice[2:4])
    fmt.Println(slice[4:])

    // Append
    slice = append(slice, 7)
    fmt.Println(slice)

    newSlice := []int{8,9,10}
    slice = append(slice, newSlice...)
    fmt.Println(slice)
}
```