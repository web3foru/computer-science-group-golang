package ListaLigadaSimple

import (
	"fmt"
)


func (n *Nodo) String() string {
	return fmt.Sprintf("Nodo(%v)", n.data)
}

func (ll *ListaLigada) EstaVacio() bool {
	return ll.nodoCabeza == nil
}

func NuevaLista() *ListaLigada {
	return &ListaLigada{nodoCabeza: nil, longitud: 0}
}

func (ll *ListaLigada) String() string {
	nodos := ""
	actual := ll.nodoCabeza
	for actual != nil {
		nodos += fmt.Sprintf("%v -> ", actual)
		actual = actual.siguiente
	}
	nodos += "nil"
	return fmt.Sprintf("ListaLigada(%v)", nodos)
}


func (ll *ListaLigada) AgregarFinal(data interface{}){
	nuevoNodo := &Nodo{data: data, siguiente: nil}
	if ll.EstaVacio() {
		ll.nodoCabeza = nuevoNodo
		fmt.Printf("Agregado al final %s, como cabeza\n", nuevoNodo)
	} else {
		actual := ll.nodoCabeza
		for actual.siguiente != nil {
			actual = actual.siguiente
		}
		actual.siguiente = nuevoNodo
		fmt.Printf("Agregado al final %s\n", nuevoNodo)
	}
	ll.longitud++
}

func (ll *ListaLigada) AgregarInicio(data interface{}){
	nuevoNodo := &Nodo{data: data, siguiente: nil}
	nuevoNodo.siguiente = ll.nodoCabeza
	ll.nodoCabeza = nuevoNodo
	ll.longitud++
	fmt.Println("Agregado al inicio %s", nuevoNodo)
}

func (ll *ListaLigada) InsertarLuegoDe(dataObjetivo interface{}, data interface{}) bool{
	nodoActual := ll.nodoCabeza
	for nodoActual != nil {
		if nodoActual.data == dataObjetivo {
			nuevoNodo := &Nodo{data: data, siguiente: nil}
			nuevoNodo.siguiente = nodoActual.siguiente
			nodoActual.siguiente = nuevoNodo
			ll.longitud++
			fmt.Println("Insertado %s luego de %s", nuevoNodo, nodoActual)
			return true
		}
		nodoActual = nodoActual.siguiente
	}
	return false
}

func (ll *ListaLigada) Eliminar(data interface{}) bool {
	nodoActual := ll.nodoCabeza
	var nodoAnterior *Nodo = nil
	for nodoActual != nil {
		if nodoActual.data == data {
			if nodoAnterior == nil {
				nodoAnterior.siguiente = nodoActual.siguiente
				fmt.Println("Eliminado %s", nodoActual)
			} else {
				ll.nodoCabeza = nodoActual.siguiente
				fmt.Println("Eliminado %s", nodoActual)
			}
			ll.longitud--
			return true
		}
		nodoAnterior = nodoActual
		nodoActual = nodoActual.siguiente
	}
	fmt.Println("No se encontró %s", data)
	return false
}

