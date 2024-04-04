package main

import (
	"fmt"
	"runtime"
	"github.com/emic310/godesde0/variables"
)

func main() {
	// fmt.Println("Hola Mundo")
	// variables.MuestroEnteros()
	// variables.RestoVariables()
	
	// declaro las variables de manera asignativa separadas por coma
	/* estado, texto := variables.ConviertoATexto(1567)

	fmt.Println("pudo convertir = ", estado)
	fmt.Println("texto convertido", texto) 
	*/

	// el if no lleva parentesis, se puede asignar antes de evaluar separando con ;
	// la llave { va en la misma linea (no abajo) y el ELSE va en la misma linea donde cierra el IF (no abajo)
	if os:= runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows, es " os)
	} else {
		fmt.Println("Esto es Windows")
	}

}
