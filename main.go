package main

import (
	"fmt"

	"github.com/robertobouses/banco5/movimientos"
)

func main() {
	var monto int
	var opcion string
	dato := movimientos.Cliente{}
	valor := movimientos.CuentaBancaria{}

	InventarioClientes := dato.FormularioCliente()
	InventarioCuentaBancarias = valor.FormularioCuenta(movimientos.Dni, InventarioCuentaBancarias)

	for {

		fmt.Println("")
		fmt.Println("Desea seguir haciendo gestiones?")
		fmt.Scanln(&opcion)
		if opcion == "NO" {
			break
		}
		fmt.Println("Iniciando GESTOR de MOVIMIENTOS")
		fmt.Println("indica el DNI del CLIENTE")
		fmt.Scanln(&movimientos.Dni)
		fmt.Println("accediendo a la cuenta de", InventarioClientes[movimientos.Dni])
		fmt.Println("indica monto de la operación")
		fmt.Scanln(&monto)

		fmt.Println("")
		fmt.Println("indica opción:")
		fmt.Println("CARGO")
		fmt.Println("ABONO")
		fmt.Println("SALIR")
		fmt.Println("")
		fmt.Scanln(&opcion)

		switch opcion {
		case "CARGO":
			InventarioCuentaBancarias := valor.Cargo(monto, InventarioClientes, movimientos.Dni, movimientos.InventarioCuentaBancarias)
			fmt.Println("Inventario cuentaaaaa", InventarioCuentaBancarias)
		case "ABONO":
			valor.Abono(monto, InventarioClientes, movimientos.Dni)

			//case "GRABAR":
			//	InventarioClientes := dato.Formulario()

		case "SALIR":
			break
		default:
			fmt.Println("Indique una opción existente")

		}
	}

	fmt.Println("Desconexión")

}
