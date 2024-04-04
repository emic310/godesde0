package funciones

import "fmt"

// closure, la ventaja es el ofuscamiento del código 

// tabla retorna una función anónima que devuelve un entero
func tabla(valor int) func() int {

	// se inicializan en la primera ejecución y luego quedan ocultas
	numero := valor
	secuencia := 0
	
	// lo que retorno es esto
	// retorno una función anónima que devuelve un entero
	return func() int {
		secuencia++
		return numero*secuencia
	}

}

func LlamarClosure() {

	tablaDel := 2
	// esta variable recibe el closure
	MiTabla := tabla(tablaDel)

	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
		// cada vez que llamo MiTabla (que tiene el closure) estoy llamando a una función 
		// se va a ejecutar lo de adentro del return de tabla (la inicialización de variables se hace solo la primera vez) 
	}

}
