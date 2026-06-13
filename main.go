package main

import (
	"errors"
	"fmt"
)

type OperacionEcommerse interface {
	Procesar() error
	MostrarDetalles()
}

type Producto struct {
	id     int
	nombre string
	precio float64
	stock  int
}

func NuevoProducto(id int, nombre string, precio float64, stock int) *Producto {
	return &Producto{
		id:     id,
		nombre: nombre,
		precio: precio,
		stock:  stock,
	}
}

func (p *Producto) ObtenerNombre() string {
	return p.nombre
}

func (p *Producto) ObtenerStock() int {
	return p.stock
}

func (p *Producto) DescontarStock(cant int) error {
	if cant <= 0 {
		return errors.New("cantidad invalida")
	}
	if p.stock < cant {
		return fmt.Errorf("no hay stock de %s, actual: %d", p.nombre, p.stock)
	}
	p.stock -= cant
	return nil
}

func (p *Producto) Procesar() error {
	fmt.Printf("Revisando stock de: %s...\n", p.nombre)
	if p.stock < 3 {
		fmt.Println("Alerta: Stock bajo minimos")
	}
	return nil
}

func (p *Producto) MostrarDetalles() {
	fmt.Printf("Item: %s | Precio: %.2f | Quedan: %d\n", p.nombre, p.precio, p.stock)
}

func main() {
	fmt.Println("=== AVANCE AUTONOMO 2 ===")

	p := NuevoProducto(101, "Laptop Core i7", 850.00, 5)

	var op OperacionEcommerse = p

	fmt.Println("Datos:")
	op.MostrarDetalles()

	fmt.Println("\nComprando 2...")
	err := p.DescontarStock(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ok")
	}
	op.MostrarDetalles()

	fmt.Println("\nComprando 5 mas...")
	err = p.DescontarStock(5)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nInterfaz:")
	op.Procesar()
}
