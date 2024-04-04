package funciones

import "fmt"

func Exponencial(valor int) {
	
	if valor > 100000000 {
		return
	}

	fmt.Println(valor)

	Exponencial(valor*2)

}

func ExponencialMia(valor int) int {
	
	fmt.Println(valor)
	
	if valor < 100000000 {
		return Exponencial(valor*2)
	}

	return valor

}

func ExponencialMiaImprimeInverso(valor int) int {
	
	if valor < 100000000 {
		return Exponencial(valor*2)
	}

	fmt.Println(valor)
	
	return valor
	
}