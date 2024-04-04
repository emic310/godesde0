package variables

import (
	"fmt"
	"time"
)

func alcance() {
	// al estar en el mismo paquete tengo acceso a todo lo publico de los otros modulos de este paquete
	MuestroEnteros()
}

// tambien pasa con las variables, una variable publica, que este fuera del metodo, 
// es vista desde todo el paquete y desde todos los lugares desde donde importen el paquete

var Nombre string
// al crear un string por declaracion, se define con el valor minimo, una cadena vacia

var Estado bool
// bool por declaracion es false???

var Sueldo float32
var Fecha time.Time

func RestoVariables() {

	fmt.Println("Estado antes de asignarle valor", Estado)

	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)

}