
package models

type Pedido struct {
	ID             int
	ClienteID      int
	ProductoID     int
	NombreProducto string
	Cantidad       int
	Total          float64
}
