package variables

import "fmt"

func MuestroEnteros() {
	var intcomun int
	intde32 := int32(1000)
	intde64 := int64(100000)

	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}
