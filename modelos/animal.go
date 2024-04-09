package modelos

type Animal struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (a *Animal) Respirar()    		{ a.respirando = true }
func (a *Animal) Comer()       		{ a.comiendo = true }
func (a *Animal) EsCarnivoro() bool { 
	a.carnivoro = true 
	return a.carnivoro
}

func (a *Animal) EstaVivo() bool { 
	a.vivo = true
	return a.vivo 
}
