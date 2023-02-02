package movimientos

import "fmt"

type CuentaBancaria struct {
	movimiento int
	saldo      int
}

var InventarioCuentaBancarias map[int]CuentaBancaria

func (c CuentaBancaria) Cargo(monto int, InventarioClientes map[int]Cliente, Dni int, InventarioCuentaBancaria map[int]CuentaBancaria) map[int]CuentaBancaria {

	InventarioCuentaBancarias = make(map[int]CuentaBancaria)
	c.saldo += monto
	fmt.Printf("La cuenta de %s tiene un saldo de %v\n", InventarioClientes[Dni].Nombre, c.saldo)

	CuentaBancaria1 := CuentaBancaria{
		saldo: c.saldo,
	}

	InventarioCuentaBancarias[Dni] = CuentaBancaria1

	return InventarioCuentaBancarias
}
