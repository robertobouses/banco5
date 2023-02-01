package movimientos

import "fmt"

var Dni int
var opcion string

type Cliente struct {
	Nombre   string
	Apellido string
}

func (c *Cliente) Formulario() map[int]Cliente {

	var InventarioClientes map[int]Cliente
	InventarioClientes = make(map[int]Cliente)

	fmt.Println("FORUMULARIO DE ADMISIÓN DE DATOS")
	for {
		fmt.Println("Indica el DNI, el Nombre y el Apellido del Cliente")
		fmt.Scanln(&Dni)
		fmt.Scanln(&c.Nombre)
		fmt.Scanln(&c.Apellido)

		Cliente1 := Cliente{

			Nombre:   c.Nombre,
			Apellido: c.Apellido,
		}

		InventarioClientes[Dni] = Cliente1

		fmt.Println("Quiere grabar otro CLIENT?")
		fmt.Scanln(&opcion)
		if opcion == "NO" {
			break
		}

	}
	return InventarioClientes
}
