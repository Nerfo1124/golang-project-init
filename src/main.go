package main

import (
	"fmt"
	"math"
)

type carPublic struct {
	brand string
	year  int
}

func (c carPublic) String() string {
	return fmt.Sprintf("Marca Carro: %s - Modelo: %d", c.brand, c.year)
}

func areaCirculoFunc(radio float64) float64 {
	return math.Pi * radio * radio
}
func areaRectanguloFunc(base float64, altura float64) float64 {
	return base * altura
}

func areaTrapezoideFunc(B float64, b float64, h float64) float64 {
	return h * (B + b) / 2
}

func contadorVocales(palabra string) (int, int, int, int, int) {
	conta := 0
	conte := 0
	conti := 0
	conto := 0
	contu := 0
	for _, valor := range palabra {
		variable := string(valor)
		switch variable {
		case "a":
			conta++
		case "e":
			conte++
		case "i":
			conti++
		case "o":
			conto++
		case "u":
			contu++
		}
	}
	return conta, conte, conti, conto, contu
}

func main() {
	fmt.Println("Hola Mundo")
	textoPrueba := "Nerfo Test"

	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.16

	fmt.Println("Pi:", pi)
	fmt.Println("Pi2:", pi2)

	//Declaracion de varaibles
	base := 12          //Primera forma
	var altura int = 14 //Segunda forma
	var area int        //Go no compila si las variables no son usadas

	fmt.Println(base, altura, area)

	//Zero values
	//Go asigna vaalores a variables vacías
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Ejercicio
	//Calcular el áera del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("El área del cuadrado es:", areaCuadrado)

	// Rectángulo
	baseRectangulo := 20
	alturaRectangulo := 10

	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println("El Area del Rectángulo es :", areaRectangulo)

	// Circulo : AreaCirculo = pi por radio al cudrado
	const PI float64 = 3.14 // Constant
	var radioCirculo float64 = 10

	areaCirculo := PI * radioCirculo * radioCirculo
	fmt.Println("El Area del Circulo es :", areaCirculo)

	// Trapecio
	var baseUno float64 = 6
	var baseDos float64 = 15
	var alturaTrapecio float64 = 25

	areaTrapecio := ((baseUno + baseDos) * alturaTrapecio) / 2
	fmt.Println("El Area del Trapecio es :", areaTrapecio)

	fmt.Printf("Circulo %.2f \n", areaCirculoFunc(2))
	fmt.Printf("Rectangulo %.2f \n", areaRectanguloFunc(5, 10))
	fmt.Printf("Trapezoide %.2f \n", areaTrapezoideFunc(10, 5, 3))

	palabra := "ejemplo para contar vocales en una frase o palabra"
	a, e, i, o, u := contadorVocales(palabra)
	fmt.Printf("la frase '%s' tiene: \n", palabra)
	fmt.Printf("%d vocales a \n", a)
	fmt.Printf("%d vocales e \n", e)
	fmt.Printf("%d vocales i \n", i)
	fmt.Printf("%d vocales o \n", o)
	fmt.Printf("%d vocales u \n", u)

	myCar := carPublic{brand: "Ford", year: 2023}
	fmt.Println(myCar)

	var otherCar carPublic
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
