package ejercicios

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func TablaCeroAlDiez() {
	
	var numero int
	var err error
	var salida string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese un número entero : ")
	if scanner.Scan() {
		
		for numero, err = strconv.Atoi(scanner.Text()); err != nil; numero, err = strconv.Atoi(scanner.Text()) {
		
			fmt.Println("Lo ingresado no es un número entero, intente nuevamente: ")
			if scanner.Scan() {
				continue
			}

		}
		
	}

	fmt.Println("Tabla del ", numero)

	for i := 0; i <= 10; i++ {
		salida = strconv.Itoa(numero) + " * " + strconv.Itoa(i) + " = "
		fmt.Println(salida, numero*i)
	}

}

// la solución del profesor
func TablaDeMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin)

	var numero int
	var err error

	for {
		fmt.Println("Ingrese un número : ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		// %d número decimal de tipo base 10
		fmt.Printf("%d x %d = %d \n", numero, i, numero*i)
	}

}
