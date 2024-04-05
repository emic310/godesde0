package mapas

import "fmt"

func MostrarMapas() {
	// el mapa se crea con el make, pero se usa la palabra reservada map, tambien le puedo dar un tamaño si quisiera pero no es necesario
	paises := make(map[string]string)
	//	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"

	/*	
	// si paso el map me imprime -> map[clave:valor clave:valor ...]
	fmt.Println(paises)
	// si paso una clave del map, me imprime solo el valor
	fmt.Println(paises["Argentina"]) 
	*/

	// creando un map por asignación directa
	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	fmt.Println(campeonato)

	// los mapas no se imprimen ordenados por clave ni valor (con los for)

	// si usamos el for range (tipo for each), me devuelve clave y valor, entonces tengo que declarar 2 variables
	/*	for equipo, puntaje := range campeonato {
			fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
		}
	*/

	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
	// %t indica que es un bool
}