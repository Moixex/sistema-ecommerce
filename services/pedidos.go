
package services

import "sistema-ecommerce/models"

func RegistrarPedido(pedidos []models.Pedido, pedido models.Pedido) []models.Pedido {
	return append(pedidos, pedido)
}

func ListarPedidos(pedidos []models.Pedido) []models.Pedido {
	return pedidos
}
