package ejercicios

import (
	"strconv"
)

func ConvertStringToInt (texto string) (int, string) {

	if entero := strconv.Atoi(texto); entero > 100 {
		return entero, "es mayor a 100"
	}

	return entero, "es menor o igual a 100"

}
