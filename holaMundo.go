package triagulo

import (
	"fmt"
	"math"
)

func triagulo() {

	fmt.Println("hola mundo")

	var nombre string
	fmt.Println("cual es tu nombre??")
	fmt.Scan(&nombre)

	fmt.Println("Hola", nombre, "bienvenido al curso de GoLang")

	var longitud1, longitud2 float64

	fmt.Println("este programa solicita las longitududes de triangulo redtangulo")
	fmt.Println("longitud 1")
	fmt.Scan(&longitud1)
	fmt.Println("longitud 2")
	fmt.Scan(&longitud2)

	area := longitud1 * longitud2 / 2
	c := math.Sqrt(longitud1*longitud1 + longitud2*longitud2)
	perimeter := (longitud1 + longitud2 + c)

	fmt.Println("el area del triangulo rectangulo es:", area)
	fmt.Println("el perimetro del triangulo rectangulo es:", perimeter)

}
