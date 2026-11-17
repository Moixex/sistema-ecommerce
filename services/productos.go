
package services

import "sistema-ecommerce/models"

func RegistrarProducto(productos []models.Producto, producto models.Producto) []models.Producto {
	return append(productos, producto)
}

func ListarProductos(productos []models.Producto) []models.Producto {
	return productos
}

func BuscarProducto(productos []models.Producto, nombre string) (models.Producto, bool) {
	for _, producto := range productos {
		if producto.Nombre == nombre {
			return producto, true
		}
	}
	return models.Producto{}, false
}

func ActualizarInventario(productos []models.Producto, id int, stock int) ([]models.Producto, bool) {
	nuevosProductos := make([]models.Producto, len(productos))
	copy(nuevosProductos, productos)

	for i, producto := range nuevosProductos {
		if producto.ID == id && stock >= 0 {
			nuevosProductos[i].Stock = stock
			return nuevosProductos, true
		}
	}

	return productos, false
}

func CrearPedido(productos []models.Producto, id int, cantidad int) ([]models.Producto, float64, bool) {
	if cantidad <= 0 {
		return productos, 0, false
	}

	nuevosProductos := make([]models.Producto, len(productos))
	copy(nuevosProductos, productos)

	for i, producto := range nuevosProductos {
		if producto.ID == id {
			if producto.Stock < cantidad {
				return productos, 0, false
			}

			total := producto.Precio * float64(cantidad)
			nuevosProductos[i].Stock -= cantidad

			return nuevosProductos, total, true
		}
	}

	return productos, 0, false
}
