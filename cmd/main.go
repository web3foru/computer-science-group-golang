package main

import (
	"fmt"
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/lista_ligada"
)

func main() {
	// Crear una instancia de ListaLigada
	lista := lista_ligada.NuevaLista()

	// Verificar si la lista está vacía
	fmt.Println("¿La lista está vacía?", lista.EstaVacio()) // Output: true

	// Agregar nodos al final de la lista
	lista.AgregarFinal(10)
	lista.AgregarFinal(20)
	lista.AgregarFinal(30)
	fmt.Println("Lista después de agregar al final: ", lista) // Output: ListaLigada(10 -> 20 -> 30)

}
