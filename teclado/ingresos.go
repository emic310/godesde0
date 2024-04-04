package teclado

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

// fmt tiene funciones Scan para leer, pero no es tan potente como las de bufio (buffer input output) 
// y da problemas en Windows

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {

	// bufio puede leer de muchas fuente de datos (teclado, lectores de código de barra, archivos...)
	// tenemos que especificar que va a leer del teclado, que es el standar input (standar output es la consola)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese número 1 : ")
	// entra al if solo si el usuario ingreso algo por teclado
	if scanner.Scan() {
		// todo lo que entra por bufio entra como texto y entra por la propiedad Text()
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			// si realmente queremos que aborte la aplicación ante un error usamos panic
			panic("El dato ingresado no es un número " + err.Error())
		}
	}

	fmt.Println("Ingrese número 2 : ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado no es un número " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda : ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, " ", numero1 * numero2)
}
