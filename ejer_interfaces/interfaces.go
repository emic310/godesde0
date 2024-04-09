package ejer_interfaces

import (
	"fmt"

	"github.com/emic310/godesde0/interfaces"
)

// esto es polimorfismo, puedo pasarle cualquier clase que implemente la interface humano (hombre / mujer)
func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}

// esto es polimorfismo, puedo pasarle cualquier clase que implemente la interface servivo (humano (hombre / mujer), animal)
func EstaVivoEsteSerVivo(sv interfaces.SerVivo) {
	var vivo = sv.EstaVivo()
	if vivo {
		fmt.Printf("Soy un ser vivo y estoy vivo \n")
	} else {
		fmt.Printf("Soy un ser vivo pero estoy muerto \n")
	}

}
