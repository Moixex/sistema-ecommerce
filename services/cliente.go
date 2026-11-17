
package services

import "sistema-ecommerce/models"

func RegistrarCliente(clientes []models.Cliente, cliente models.Cliente) []models.Cliente {
	return append(clientes, cliente)
}

func ListarClientes(clientes []models.Cliente) []models.Cliente {
	return clientes
}

func BuscarCliente(clientes []models.Cliente, id int) (models.Cliente, bool) {
	for _, cliente := range clientes {
		if cliente.ID == id {
			return cliente, true
		}
	}

	return models.Cliente{}, false
}
