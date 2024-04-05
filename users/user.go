package users

import (
	"fmt"
	"time"
	"github.com/emic310/godesde0/modelos"
)

func AltaUsuario() {
	
	// recién acá estoy creando (instanciando) un objeto de tipo user
	u := new(modelos.User)
	
	// le pongo los datos, acá podria directamente asignarle valores a cada campo de del User con u.talCosa = x
	// pero tengo una función AddUser que se encarga de eso y podría tener validaciones u otra lógica (conectarse a BD, actualizar o rechazar si ya existe)
	u.AddUser(10, "Pablo", time.Now(), true)

	fmt.Println(u)
}