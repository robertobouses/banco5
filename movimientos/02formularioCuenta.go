package movimientos

import "fmt"

func (c CuentaBancaria) FormularioCuenta(Dni int, InventarioCliente map[int]Cliente) map[int]CuentaBancaria {
	var InventarioCuentaBancarias map[int]CuentaBancaria
	InventarioCuentaBancarias = make(map[int]CuentaBancaria)

	fmt.Println("Indique el saldo inicial de la cuenta de", InventarioCliente[Dni].Nombre)
	fmt.Scanln(&c.saldo)
	CuentaBancaria1 := CuentaBancaria{
		saldo: c.saldo,
	}

	InventarioCuentaBancarias[Dni] = CuentaBancaria1

	return InventarioCuentaBancarias
}
