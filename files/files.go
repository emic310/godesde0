package files

import (
	//"bufio"
	"fmt"
	"github.com/emic310/godesde0/ejercicios"
	//"ioutil"
	"os"
)

// nombre del archivo tabla.txt, hay que darle la url desde la raiz 
// porque lo voy a estar ejecutando desde main que esta en la raiz
var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	
	var texto string = ejercicios.TablaDeMultiplicar()

	// crea un archivo con el nombre que definimos, si el archivo ya existe lo elimina y crea uno nuevo
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		// si falla que salga del metodo 
		return
	}

	// si creo el texto, lo guardo en el archivo
	fmt.Fprintln(archivo, texto)

	// siempre que se crea, abre para lextura, escritura o se lee  hay que cerrarlo
	archivo.Close()

}

func SumaTabla() {
	
	var texto string = ejercicios.TablaDeMultiplicar()

	// si no pudo concatenar tiro un mensaje
	if !Append(texto) {
		ftm.Println("Error al concatenar contenido")
	}

}

func Append(texto string) bool {
	// archivo es un buffer que me crea el SO
	
	// os.Open() abre el archivo en modo solo para leectura
	// os.OpenFile() abre el archivo en modo Append, me permite abrir y concatenarle algo al archivo
	
	// le pasamos 2 flags para que sepa que vamos a escribir y agregar
	// con el | podemos pasar mas de un flag aunque el argumento que pide sea solo uno
	// le pasasmos escritura y append, os.O_WRONLY|os.O_APPEND
	// si no le pasamos el flag Append, me lo abre de de escritura pero va a limpiar el contenido del archivo y poner solo lo nuevo

	// luego tambien me pide los permisos -> owner lectura/escritura (6), para grupo y usuarios externos solo lectura (4)

	// aca solo abro el archivo con los modos y permisos necesarios pero no graba nada aun
	archivo, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		// como fallo retorno false
		return false
	}

	// aca si no falló, grabo los datos
	// el primer valor que me retorna lo ignoro _
	// esta funcion escribe en el archivo (buffer) lo que le paso en texto (parametro)
	_, err = archivo.WriteString("/n" + texto)
	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		// como fallo retorno false
		return false
	}

	// si no falló, ya se grabó, entonces cerramos el archivo
	archivo.Close()
	
	return true

}
