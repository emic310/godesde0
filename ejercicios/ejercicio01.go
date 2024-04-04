package ejercicios

import (
	"strconv"
)

func ConvertStringToInt(texto string) (int, string) {
	// si no me interesa una de las variables que retorna, puedo "descartarla" con _
	// entero, _ := strconv.Atoi(texto)
	
	entero, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "no se pudo convertir " + err.Error()
	}

	if entero > 100 {
		return entero, "es mayor a 100"
	}

	return entero, "es menor o igual a 100"
}