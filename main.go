package main

import (
	"errors"
	"fmt"
)

// Funcion classic
func suma(a, b int) int{
	return a + b
}

// Funcion que devuelve mas de un valor
func dividir(a, b float64) (float64, error){

	if (b == 0){
		return 0 , errors.New("no se puede dividir por 0")
	}

	cociente := a / b
	// resto := float64(int(a) % int(b))

	return cociente, nil
}

// Funcion con numero variable de argumentos
func imprimirNombres(nombres ...string){
	for _, nombre := range nombres{
		fmt.Println(nombre)
	}
}

// Ejmemplo de closure
var a = 0
func contador() func() int{
	c := 0

	return func() int {
		c++
		return c + a
	}
}

type Rectangulo struct{
	Ancho, Alto float64
}

func(r Rectangulo) Area() float64{
	return r.Ancho * r.Alto
}


func main() {
	// cociente, error := dividir(10,0)

	// if error != nil {
	// 	// comprobaciones
	// 	fmt.Println(error)
	// }

	// fmt.Println(cociente)

	
	// Closure
	cont := contador()
	fmt.Println("contador: ", cont())
	fmt.Println("contador: ", cont())
	fmt.Println("contador: ", cont())
	fmt.Println("contador: ", cont())
	fmt.Println("contador: ", cont())
	fmt.Println("contador: ", cont())
	fmt.Println("contador: ", cont())
	fmt.Println("contador: ", cont())
	fmt.Println("contador: ", cont())
	fmt.Println("contador: ", cont())
	fmt.Println("contador: ", cont())

	rect := Rectangulo{Ancho: 10, Alto: 5}
	fmt.Println("Area del rectangulo: ", rect.Area())
}
