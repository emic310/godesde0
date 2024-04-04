package ejercicios

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

var numero1 int
var err error
var salida string

func TablaCeroAlDiez() {
	
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese un número entero : ")
	if scanner.Scan() {
		
		for numero1, err = strconv.Atoi(scanner.Text()); err != nil; numero1, err = strconv.Atoi(scanner.Text()){
		
			fmt.Println("Lo ingresado no es un número entero, intente nuevamente: ")
			if scanner.Scan() {
				continue
			}

		}
		
	}

	fmt.Println("Tabla del ", numero1)

	for i := 0; i <= 10; i++ {
		salida = strconv.Itoa(numero1) + " * " + strconv.Itoa(i) + " = "
		fmt.Println(salida, numero1 * i)
	}

}