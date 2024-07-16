package users

import (
	"fmt"
	"time"
	"github.com/danar37/godesde0/modelos"

)


func AltaUsuario(){
	u := new(modelos.User)
	u.AddUser(10, "Daniel", time.Now(), true)
	fmt.Println(u)

}