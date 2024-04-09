package interfaces

type Vegetal interface {
	//funciones propias del vegetal
	ClasificacionVegetal() string
	// implementa la interface SerVivo
	Vivir()
	EstaVivo() bool
}
