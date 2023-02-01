package movimientos

import "fmt"

type CuentaBancaria struct {
	movimiento int
	saldo      int
}

func (c CuentaBancaria) Cargo(monto int, InventarioClientes map[int]Cliente, Dni int) {

	c.saldo += monto
	fmt.Println("El saldo es", InventarioClientes[Dni].Nombre, c.saldo)

}
