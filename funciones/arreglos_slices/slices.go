package arreglos_slices

import "fmt"

// no puedo volver a usar el nombre tabla porque ya hay una variable con ese nombre en el paquete (en el modulo arreglos.go)
// el Slice (es un arreglo dinámico) se declara igual pero sin darlo un largo al arreglo []int
var otraTabla []int = []int{2, 5, 4}

// arreglo común
var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

// arreglo tiene -> [6, 98, 21, 674, 18, 36, 78, 9, 0, 0]

func MuestroSlice() {
	fmt.Println(otraTabla)

	// me permite crear Slices a partir de arreglos como modelo (tipo de dato, elementos que tiene)
	todo := arreglo[:] // Slice creado desde un vector, desde el inicio hasta el final (Slice con mismos datos del arreglo)

	// y permite tomar algunos elementos si asi lo deseo
	porcion := arreglo[3:]   // Slice creado desde un vector, desde la posición 3
	porcion2 := arreglo[:5]  // Slice creado desde un vector, desde la posición 0 hasta la 5
	porcion3 := arreglo[6:7] // Slice creado desde un vector, desde la posición 6 hasta la 7

	fmt.Println(todo)
	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

