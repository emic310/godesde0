package interfaces

// al igual que las estructuras, las interfaces empiezan con type
type Humano interface {
	//funciones propias del humano (podria tener una interface para las cosas en comun de humano y animal)
	Respirar()
	Comer()
	Pensar()
	Sexo() string
	// implementa la interface SerVivo
	Vivir()
	EstaVivo() bool
}
