package ejercicios

import "strconv"

func PrimeraFuncion(texto string) (int, string) {
	conv, myerror := strconv.Atoi(texto)
	if myerror != nil {
		return 0, "Hubo un error por acá " + myerror.Error()
	}

	if conv > 100 {
		return conv, "Es mayor a 100"
	} else {
		return conv, "Es menor a 100"
	}

}
