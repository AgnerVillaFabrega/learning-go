package main

import (
	"fmt"
	"strings"
)

func main() {
	// Numeros
	// var entero int = 10
	entero := 10
	decimal := 3.14
	suma := entero + int(decimal)

	// Textos
	mensaje := "Hola,"
	concatenado := mensaje + "Gentleman"
	enMayusculas := strings.ToUpper(concatenado)

	// Booleanos
	esVerdadero := true

	// Arrays y Slices
	arrayFijo := [3]int{1,2,3}
	sliceVariable := []int{4,5,6}
	sliceVariable = append(sliceVariable, 7)

	// Mapas
	diccionario := map[string]int{
		"clave1" : 1,
		"clave2": 2,
	}

	// Structs
	type Persona struct {
		Nombre string
		Edad int
	}

	persona := Persona{Nombre: "Agner", Edad: 22}


	// Imprimir resultados
	fmt.Println("Suma", suma)
	fmt.Println("Mensaje", enMayusculas)
	fmt.Println("Array", arrayFijo)
	fmt.Println("Slice", sliceVariable)
	fmt.Println("Mapa", diccionario)
	fmt.Println("Struct", persona)
	fmt.Println("Booleano", esVerdadero)



}

//? Tipos
// bool [true/false]-> falg o condicionales -> false
// string [Cadena de caracteres]-> representar texto -> ""
// int, int8, int16, int32, int64 [entero]-> controlar el tamano de los enteros -> 0
// uint ,uint8,uint16,uint32,uint64 [entero sin signo]-> cuando no queremos negativos ->0
// float32, float64 [Representar valores numericos reales]-> numeros decimales => 32|64 el sistema -> 0
// byte === uint8 -> trabajar con datos binarios -> 0
// rune === int32 -> representar un solo caracter que ocupa mas de un byte 🏆 -> 0
// complex64, complex12 [Cuando tiene una parte real y otra imaginaria]-> N + iN -> 0 +
