package modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

// funciones de la interface humano (no se pone el implements al definir la clase (struct) )
// ya al agregar las mismas funciones que tiene definida la interface, detecta que esta clase implementa esa interface
func (h *Hombre) Respirar()    { h.respirando = true }
func (h *Hombre) Comer()       { h.comiendo = true }
func (h *Hombre) Pensar()      { h.pensando = true }
func (h *Hombre) Sexo() string { return "Hombre" }

// funcion de la interface humano pero que es heredada de la interface servivo
func (h *Hombre) Vivir()         { h.vivo = true }
func (h *Hombre) EstaVivo() bool { return h.vivo }
