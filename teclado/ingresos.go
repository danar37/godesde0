package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var dan1 int
var dan2 int
var ley string
var myerror error

func IngresarconTeclado() {
	key := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese dan1 : ")
	if key.Scan() {
		dan1, myerror = strconv.Atoi(key.Text())
		if myerror != nil {
			panic("Colocaste cualquier cosa " + myerror.Error())
		}

	}

	fmt.Println("Ingrese dan2 : ")
	if key.Scan() {
		dan2, myerror = strconv.Atoi(key.Text())
		if myerror != nil {
			panic("Colocaste cualquier cosa " + myerror.Error())
		}

	}

	fmt.Println("Ingrese ley : ")
	if key.Scan() {
		ley = key.Text()
	}

	fmt.Println(ley, dan1+dan2)

}
