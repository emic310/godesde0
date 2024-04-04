package funciones

import "fmt"

func Calculos () {
	
	var numero3 int = 32
	var numero4 int = 243

	// asigno una función anónima a una variable
	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	fmt.Println(calculo(10, 25))

	// puedo cambiar la logica de la funcion, no su estructura definida (cambiar parametros, que retorna) 
	calculo := func(numero1 int, numero2 int, numero3 int) int {
		return numero1 * numero2 * numero3
	}

	// al poner suma me la muestra como si fuera una funcion fuera del modulo
	fmt.Println(calculo(10, 25))

}