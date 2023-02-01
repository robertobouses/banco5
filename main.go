package main

import (
	"fmt"

	"github.com/robertobouses/banco5/movimientos"
)

func main() {
	valor := movimientos.Cliente{}
	valor.Formulario()
	fmt.Println("Desconexión")
}
