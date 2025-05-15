package main

import "fmt"

// Punteros
// * -> puntero
// & -> refenrencia | Direccion de memoria
// COPIA LOS ARGUMENTOS

func incrementar(numero *int){
	// Nunca modifiques el argumento directamente
	*numero++
}

//? Porque usar punteros?
// Queremos modificar el elemento original!
// 

func main() {
	valor := 10
	fmt.Println("Valor antes de incrementar: ", valor)
	incrementar(&valor)
	fmt.Println("Valor despues de incrementar: ", valor)
	
	// new()
	puntero := new(int) //Puntero inicializado en 0
	fmt.Println("Valor inicial con new: ", *puntero)
	*puntero = 20
	fmt.Println("Valor inicial con new: ", *puntero)
}
