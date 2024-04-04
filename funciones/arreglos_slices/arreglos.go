package arreglos_slices

import "fmt"

// aca declaro y asigno valores (tengo que repetir el [10]int)
// en este caso asigno valores a los primeros 3 lugares y el resto se completa con Ceros (no hay nulos)
var tabla [10]int = [10]int{10, 0, 59}

// aca solo declaro
var matriz [20][30]int

func MuestroArreglos() {
	// aca asigno valores a posiciones especificas
	tabla[7] = 33
	tabla[2] = 54

	// creo la tabla de manera asignativa
	tabla2 := [10]string{"Pablo", "Juan"}
	
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][24] = 15
	fmt.Println(matriz)
}