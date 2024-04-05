package modelos

// en el paquete modelos solamente vamos a colocar definiciones y estructuras
// es un estandar en los desarrolos de GO, se tienen todas las definiciones y estructuras en un solo lugar
// algunas funciones propias de esta estructura si las vamos a poner acá, pero la lógica va a estar en otro paquete

import (
	"time"
)

// las estructuras empiezan con type
type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

// esta es solo la definición!!! no es el objeto

// todas las funciones tienen que hacer referencia a la estructura User !!!

// pongo a que estructura estoy referenciando entre el func y el nombre del metodo -> (usuario *User)
// va un puntero porque no se referencia al nombre User sino a la dirección de memoria
// (no se puede usar this o self en la nueva version de GO)
// si no se pone el (usuario *User) es como que la función no existe para el objeto, no hay forma de acceder a ella
// solo con models.AddUser,  pero no queda claro a que o quien le estas agregando los datos

func (usuario *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.CreatedAt = createdAt
	usuario.Status = status
}
// este metodo no es el constructor sólo cumple la función de ponerle todos los datos a la estructura