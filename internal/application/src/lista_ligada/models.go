package lista_ligada

type Nodo struct {
	data      interface{}
	siguiente *Nodo
}

type ListaLigada struct {
	nodoCabeza *Nodo
	longitud   int
}
