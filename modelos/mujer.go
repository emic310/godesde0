package modelos

type Mujer struct {
	// esto es herencia, en vez de poner todas las propiedades que tengo en hombre y tambien las tiene la mujer
	// digo que mujer hereda de hombre
	Hombre
	// y luego puedo agregar propiedades que son solo de las mujeres
	EsMadre bool
}

// funciones de la interface humano (no se pone el implements al definir la clase (struct) )
// ya al agregar las mismas funciones que tiene definida la interface, detecta que esta clase implementa esa interface
func (m *Mujer) Respirar()    { m.respirando = true }
func (m *Mujer) Comer()       { m.comiendo = true }
func (m *Mujer) Pensar()      { m.pensando = true }
func (m *Mujer) Sexo() string { return "Mujer" }

func (m *Mujer) EstaVivo() bool  { 
	h.vivo = true
	return h.vivo 
}

// funciones propizs de la clase
func (m *Mujer) EsMadre bool  { return true }