package variables

import "fmt"

// al empezar con mayusculas la funcion le estoy indicando que es publica 
func MuestroEnteros() {

	// manera declarativa, no se define en null, para int se define en Cero
	var intComun int

	// manera asignativa, le asigno lo que retorna la funcion int32(10)
	// o sea que va a tener el valor y tipo de dato que retorne la funcion
	intDe32 := int32(10)
	intDe64 := int64(100)

	// no es lo mismo int, que int32 o int64

	fmt.Println("intComun = ", intComun)
	fmt.Println("intDe32 = ", intDe32)
	fmt.Println("intDe64 = ", intDe64)
}
