package mapas

import "fmt"

func MostarMapas() {
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"] = "CDMX"
	paises["Venezuela"] = "Caracas"
	fmt.Println(paises)
	fmt.Println(paises["Caracas"])

	campeonato := map[string]int{
		"Madrid":  39,
		"Israel":  40,
		"Suiza":   41,
		"Maracay": 42,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	delete(campeonato, "Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Maracay"]

	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
}
