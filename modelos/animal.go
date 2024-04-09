package modelos

type Animal struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

// funciones de la interface servivo
func (a *Animal) Vivir()         { a.vivo = true }
func (a *Animal) EstaVivo() bool { return a.vivo }

// funciones propias de la clase
func (a *Animal) Respirar() { a.respirando = true }
func (a *Animal) Comer()    { a.comiendo = true }
func (a *Animal) EsCarnivoro() bool {
	a.carnivoro = true
	return a.carnivoro
}
