package iteraciones

import (
	"fmt"
)

func Iterar () {
// no existe while en GO

	i := 0
	for i < 3 {
		fmt.Println (i)
		i++
	}

	// forma compacta, como el clásico, declaro; condición; algo que hago 
	// previo a la siguiente iteración/evaluación de la condición 
	for i := 0; i < 5; i++ {
		fmt.Println (i)
	}

	// puedo hacer i := 0, porque i queda dentro del scope del for
	for i := 0; i < 50; i += 5 {
		fmt.Println (i)
	}

	for i := 100; i > 50; i -= 5 {
		fmt.Println (i)
	}

	// se puede cortar con break
	for i := 10; i > 0; i-- {
		if i == 6 {
			break
		}
		fmt.Println (i)
	}

	// se puede saltar lo que ejecuta dentro del for con continue, esto hace que vaya de nuevo a evaluar la condición 
	// hace i-- y luego i > 0 pero no el println
	for i := 10; i > 0; i-- {
		if i == 6 {
			continue
		}
		fmt.Println (i)
	}

}