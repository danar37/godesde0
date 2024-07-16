package main

import (
	e "github.com/danar37/godesde0/ejer_interfaces"
	"github.com/danar37/godesde0/modelos"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(2024)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Esto no es WIndows, es ", os)
	} else {
		fmt.Println("Esto es WIndows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	conv, texto := ejercicios.PrimeraFuncion("300")
	fmt.Println(conv)
	fmt.Println(texto)

	teclado.IngresarconTeclado()

	iteraciones.Iterar()

	fmt.Println(ejercicios.TablaDeMultiplicar())

	files.GrabaTabla()

	files.SumaTabla()

	files.LeoArchivo()

	funciones.Calculos()

	funciones.LlamaClosures()

	funciones.Exponencia(1)

	arreglos_slices.MuestroArreglos()

	arreglos_slices.MuestroArreglos()

	arreglos_slices.MuestroSlice()

	arreglos_slices.Capacidad()

	mapas.MostarMapas()

	users.AltaUsuario()*/

	Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)

}
