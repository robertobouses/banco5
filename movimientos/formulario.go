package movimientos

import "fmt"

type Cliente struct {
	Dni      int
	Nombre   string
	Apellido string
}

func (c *Cliente) Formulario() {
	fmt.Println("FORUMULARIO DE ADMISIÓN DE DATOS")
	fmt.Println("Indica el DNI, el Nombre y el Apellido del Cliente")
	fmt.Scanln(&c.Dni)
	fmt.Scanln(&c.Nombre)
	fmt.Scanln(&c.Apellido)

}
