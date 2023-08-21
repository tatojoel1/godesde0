package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Joel"
	Estado = true
	Sueldo = 1000.0
	Fecha = time.Now()

	fmt.Println(Nombre, Estado, Sueldo, Fecha)
}

func ConviertoaTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
