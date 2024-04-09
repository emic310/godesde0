package interfaces

type Animal interface {
	//funciones propias del animal
	Respirar()
	Comer()
	EsCarnivoro() bool
	// implementa la interface SerVivo
	Vivir()
	EstaVivo() bool
}
