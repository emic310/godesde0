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

// uso la solución del profesor para la escritura de archivo
func TablaDeMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)

	var numero int
	var err error
	var texto string

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
		// Printf retorna un numero entero y un error, entonces no nos sirve para guardar en la variable 
		// o sirve pero tenemos que descartar el error

		// Fprintln que guarda en un archivo el resultado, podriamos usarlo  pero como vamos a manejarlo en un paquete no lo usamos
		
		// Sprintf que retorna solo el string resultado
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
		// %d número decimal de tipo base 10
	}

	return texto

}
