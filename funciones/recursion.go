package funciones

import "fmt"

func Exponencial(valor int) {

	if valor > 1000000 {
		return
	}

	fmt.Println(valor)

	Exponencial(valor * 2)

}

func ExponencialMiaSinRetorno(valor int) {

	fmt.Println(valor)

	if valor < 1000000 {
		Exponencial(valor * 2)
	}

}

func ExponencialMia(valor int) int {

	if valor < 1000000 {
		fmt.Println(valor)
		return ExponencialMia(valor * 2)
	}

	return valor

}

func ExponencialMiaImprimeInverso(valor int) int {

	if valor < 1000000 {
		ExponencialMiaImprimeInverso(valor * 2)
		fmt.Println(valor)
	}

	return valor

}
