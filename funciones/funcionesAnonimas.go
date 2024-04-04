package funciones

import "fmt"

// funciones anonimas pueden ser asignadas a variables, pueden ser enviadas por parametro

// podria devolver una funcion
//func Calculos () func (int, int) int {
func Calculos () {
	
	var numero3 int = 32
	var numero4 int = 243

	// asigno una función anónima a una variable
	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	fmt.Println(calculo(10, 25))

	// puedo cambiar la lógica de la función, no su estructura definida (cambiar parametros, que retorna) 
	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3
	}

	// al poner suma me la muestra como si fuera una función fuera del modulo
	// envio la función como parámetro 
	fmt.Println(calculo(10, 25))

}