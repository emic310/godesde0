package ejercicios

import (
	"strconv"
)

func ConvertStringToInt(texto string) (int, string) {
	entero, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "no se pudo convertir"
	}

	if entero > 100 {
		return entero, "es mayor a 100"
	}

	return entero, "es menor o igual a 100"
}