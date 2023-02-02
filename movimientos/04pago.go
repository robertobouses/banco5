package movimientos

import "fmt"

func (c CuentaBancaria) Abono(monto int, InventarioClientes map[int]Cliente, Dni int) {
	c.saldo -= monto
	fmt.Printf("La cuenta de %s tiene un saldo de %v\n", InventarioClientes[Dni].Nombre, c.saldo)

}
