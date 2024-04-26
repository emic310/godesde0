package defer_panic

import (
	"fmt"
	"log"
)

// NO hay TRY/CATCH en GO, en cambio se puede usar esto.

// el defer es una instruccion que nos permite configurar cual va a ser la última instrucción en ejecutarse cuando salga 
// de la función (útil si tenemos return en medio del codigo o si sale por un error)
// un buen uso del defer por ejemplo es cerrar un buffer/archivo inmediatamente a la apertura 
func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")
	
	// se puede tener N defer dentro del mismo código, los ejecuta en orden

	fmt.Println("Este es el segundo mensaje")
}

// el panic aborta el programa con un mensaje que lo envía a consola
// el recover me permite recuperarme de un panic
// el raro usar panic + recover, pero un ejemplo es para loguear errores de rutinas que arrojen panic, cerrar conexiones...
func EjemploPanic() {
	// el recover se usa si o si con un defer, sino el recover no tiene forma de capturar el panic
	// el recover obliga al compilador a ejecutar este código aunque haya habido un paniic que aborto el sistema

	// defer solo acepta una función, si quisieramos ejecutar mas de una hay que hacerlo con una función anónima
	defer func() {
		// hay que configurar una variable de tipo recover, si hubo un panic, reco va a ser distinto de nil
		reco := recover()
		if reco != nil {
			// el log es como hacer un print, pero siempre antepone al mensaje una fecha y hora
			log.Fatalf("ocurrió un error que generó Panic \n %v", reco) 
			// fatal es equivalente al print pero ademas aborta el sistema (call to os.Exit(1))
			// %v es un objeto de tipo variante, recibe cualquier tipo de parametro
			// al pasarle reco, pasamos la información del panic
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}