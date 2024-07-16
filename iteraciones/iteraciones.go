package iteraciones

import "fmt"

func Iterar() {
	for d := 10; d > 1; d-- {
		if d == 6 {
			break
		}
		fmt.Println(d)
	}

}
