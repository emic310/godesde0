package main

// todos los modulos empiezan con la declaración de en que paquete se encuentran

// luego van los imports
import (
	//"fmt"
	//"runtime"
	//"github.com/emic310/godesde0/variables"
	//"github.com/emic310/godesde0/ejercicios"
	//"github.com/emic310/godesde0/teclado"
	//"github.com/emic310/godesde0/iteraciones"
	//"github.com/emic310/godesde0/files"
	//"github.com/emic310/godesde0/funciones"
	//"github.com/emic310/godesde0/arreglos_slices"
	//"github.com/emic310/godesde0/mapas"
	//"github.com/emic310/godesde0/users"

	//"github.com/emic310/godesde0/modelos"
	// puedo darle un alias al modulo si quiero
	//e "github.com/emic310/godesde0/ejer_interfaces"

	d "github.com/emic310/godesde0/defer_panic"
)

// luego todo lo demas, variables, metodos, funciones
func main() {
	// fmt.Println("Hola Mundo")
	// variables.MuestroEnteros()
	// variables.RestoVariables()

	// declaro las variables de manera asignativa separadas por coma
	/* estado, texto := variables.ConviertoATexto(1567)

	fmt.Println("pudo convertir = ", estado)
	fmt.Println("texto convertido", texto)
	*/

	// el if no lleva parentesis, se puede asignar antes de evaluar separando con ;
	// la llave { va en la misma linea (no abajo) y el ELSE va en la misma linea donde cierra el IF (no abajo)
	/*
		if os := runtime.GOOS; os == "linux" || os == "darwin" {
			fmt.Println("Esto no es Windows, es ", os)
		} else {
			fmt.Println("Esto es Windows")
		}

		//tambien se puede asignar antes de evaluar separando con ;
		switch os := runtime.GOOS; os {
			case "linux":
				fmt.Println("Esto es Linux")
			case "windows":
				fmt.Println("Esto es Windows")
			default:
				fmt.Println("\n Esto es ", os)
		}
	*/

	/*
		entero, texto := ejercicios.ConvertStringToInt("99")
		fmt.Println(entero, " ", texto)

		entero, texto = ejercicios.ConvertStringToInt("500")
		fmt.Println(entero, " ", texto)

		entero, texto = ejercicios.ConvertStringToInt("texto")
		fmt.Println(entero, " ", texto)
	*/

	//teclado.IngresoNumeros()

	//iteraciones.Iterar()

	//ejercicios.TablaCeroAlDiez()

	//fmt.Println(ejercicios.TablaDeMultiplicar())
	//files.GrabaTabla()
	//files.SumaTabla()
	//files.LeoArchivo()

	//funciones.Calculos()
	//funciones.LlamarClosure()

	// funciones.Exponencial(2)
	// funciones.ExponencialMiaSinRetorno(2)
	// funciones.ExponencialMia(2)
	// funciones.ExponencialMiaImprimeInverso(2)

	//arreglos_slices.MuestroArreglos()
	//arreglos_slices.MuestroSlice()
	//arreglos_slices.Capacidad()

	//mapas.MostrarMapas()

	//users.AltaUsuario()

	/* Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)
	Maria.Vivir()
	e.EstaVivoEsteSerVivo(Maria)

	SerVivo := new(modelos.Animal)
	e.EstaVivoEsteSerVivo(SerVivo) */

	// d.VemosDefer()
	d.EjemploPanic()
	
}
