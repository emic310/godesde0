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

func Capacidad() {
	// el Slice tiene un largo y una capacidad (memoria reservada)
	// la función make permite crear un Slice (y otras estructuras) definiendo el tipo de dato, el tamaño y la capacidad
	// en este caso el Slice tiene un largo de 5 y una capacidad de 20 en memoria pero no puedo acceder a los elementos 6 a 19
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	// dificilmente se use asi pero se puede hacer (tamaño y capacidad en Cero)
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		// aca le agrego elementos al Slice, tal cual como si fuera una lista
		// tengo que indicar a que Slice le agrego el o los elementos, tambien puedo agregarle otro u otros Slices
		nums = append(nums, i)

		// si el Slice se llena, se duplica la capacidad, va reservando 2^n memoria (2, 4, 8 ... 1024 ...)
		// para el compilador, para el SO, es mas costoso ir a reservar memoria que utilizarla
		// es menos costoso tener memoria sin usar, que estar reasignando memoria a la variable (reservar y moverla)
	}

	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
