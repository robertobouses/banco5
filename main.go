package main

import (
	"fmt"

	"github.com/robertobouses/banco5/movimientos"
)

func main() {
	var monto int

	dato := movimientos.Cliente{}
	InventarioClientes := dato.Formulario()
	valor := movimientos.CuentaBancaria{}

	fmt.Println("indica el DNI del CLIENTE")
	fmt.Scanln(&movimientos.Dni)

	fmt.Println("indica monto de la operación")
	fmt.Scanln(&monto)
	valor.Cargo(monto, InventarioClientes, movimientos.Dni)

	fmt.Println("Desconexión")
}
