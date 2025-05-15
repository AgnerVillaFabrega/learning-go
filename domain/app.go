package domain

import (
	"fmt"
	"learning-go/domain/entities"
)


func App (){
	persona := entities.Persona{Nombre: "Agner", Apellido: "Villa", Edad: 22}
	// elemento := entities.TipoElemento{Nombre: "Elemento"}

	fmt.Println(persona.Saludar())
	persona.CumplirAnios()

	fmt.Println(persona)
}