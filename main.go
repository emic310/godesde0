package main

import (
	"fmt"
	"github.com/emic310/godesde0/variables"
)

func main() {
	// fmt.Println("Hola Mundo")
	// variables.MuestroEnteros()
	// variables.RestoVariables()
	
	// declaro las variables de manera asignativa separadas por coma
	estado, texto := variables.ConviertoATexto(1567)

	fmt.Println("pudo convertir = ", estado)
	fmt.Println("texto convertido", texto)
}
