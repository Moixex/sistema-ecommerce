
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"sistema-ecommerce/models"
	"sistema-ecommerce/services"
)

var productos []models.Producto
var clientes []models.Cliente
var pedidos []models.Pedido

var lector = bufio.NewReader(os.Stdin)

func leerTexto(mensaje string) string {
	fmt.Print(mensaje)
	texto, _ := lector.ReadString('\n')
	return strings.TrimSpace(texto)
}

func registrarProducto() {
	fmt.Println("\n--- REGISTRAR PRODUCTO ---")

	nombre := leerTexto("Nombre: ")
	precioTexto := leerTexto("Precio: ")
	stockTexto := leerTexto("Stock: ")

	precio, err1 := strconv.ParseFloat(precioTexto, 64)
	stock, err2 := strconv.Atoi(stockTexto)

	if nombre == "" || err1 != nil || err2 != nil || precio <= 0 || stock < 0 {
		fmt.Println("Datos incorrectos.")
		return
	}

	producto := models.Producto{
		ID:     len(productos) + 1,
		Nombre: nombre,
		Precio: precio,
		Stock:  stock,
	}

	productos = services.RegistrarProducto(productos, producto)
	fmt.Println("Producto registrado correctamente.")
}

func listarProductos() {
	fmt.Println("\n--- LISTA DE PRODUCTOS ---")

	for _, p := range services.ListarProductos(productos) {
		fmt.Printf("ID: %d | %s | $%.2f | Stock: %d\n",
			p.ID, p.Nombre, p.Precio, p.Stock)
	}

	if len(productos) == 0 {
		fmt.Println("No hay productos registrados.")
	}
}

func buscarProducto() {
	nombre := leerTexto("Nombre del producto: ")

	producto, encontrado := services.BuscarProducto(productos, nombre)

	if encontrado {
		fmt.Printf("ID: %d | %s | $%.2f | Stock: %d\n",
			producto.ID, producto.Nombre, producto.Precio, producto.Stock)
	} else {
		fmt.Println("Producto no encontrado.")
	}
}

func actualizarInventario() {
	listarProductos()

	idTexto := leerTexto("ID del producto: ")
	id, err1 := strconv.Atoi(idTexto)

	stockTexto := leerTexto("Nueva cantidad: ")
	stock, err2 := strconv.Atoi(stockTexto)

	if err1 != nil || err2 != nil || stock < 0 {
		fmt.Println("Datos inválidos.")
		return
	}

	nuevosProductos, actualizado := services.ActualizarInventario(productos, id, stock)

	if actualizado {
		productos = nuevosProductos
		fmt.Println("Inventario actualizado correctamente.")
	} else {
		fmt.Println("Producto no encontrado.")
	}
}

func registrarCliente() {
	fmt.Println("\n--- REGISTRAR CLIENTE ---")

	nombre := leerTexto("Nombre: ")
	correo := leerTexto("Correo: ")
	telefono := leerTexto("Teléfono: ")

	if nombre == "" || correo == "" || telefono == "" || !strings.Contains(correo, "@") {
		fmt.Println("Datos incorrectos.")
		return
	}

	cliente := models.Cliente{
		ID:       len(clientes) + 1,
		Nombre:   nombre,
		Correo:   correo,
		Telefono: telefono,
	}

	clientes = services.RegistrarCliente(clientes, cliente)
	fmt.Println("Cliente registrado correctamente.")
}

func listarClientes() {
	fmt.Println("\n--- LISTA DE CLIENTES ---")

	if len(clientes) == 0 {
		fmt.Println("No hay clientes registrados.")
		return
	}

	for _, c := range services.ListarClientes(clientes) {
		fmt.Printf("ID: %d | Nombre: %s | Correo: %s | Teléfono: %s\n",
			c.ID, c.Nombre, c.Correo, c.Telefono)
	}
}

func crearPedido() {
	fmt.Println("\n--- CREAR PEDIDO ---")

	if len(clientes) == 0 || len(productos) == 0 {
		fmt.Println("Debe registrar al menos un cliente y un producto.")
		return
	}

	listarClientes()
	clienteTexto := leerTexto("ID del cliente: ")
	clienteID, err := strconv.Atoi(clienteTexto)

	if err != nil {
		fmt.Println("ID inválido.")
		return
	}

	cliente, encontrado := services.BuscarCliente(clientes, clienteID)

	if !encontrado {
		fmt.Println("Cliente no encontrado.")
		return
	}

	listarProductos()

	idTexto := leerTexto("ID del producto: ")
	id, err := strconv.Atoi(idTexto)

	if err != nil {
		fmt.Println("ID inválido.")
		return
	}

	cantidadTexto := leerTexto("Cantidad: ")
	cantidad, err := strconv.Atoi(cantidadTexto)

	if err != nil || cantidad <= 0 {
		fmt.Println("Cantidad inválida.")
		return
	}

	nuevosProductos, total, realizado := services.CrearPedido(productos, id, cantidad)

	if !realizado {
		fmt.Println("No se pudo realizar el pedido.")
		return
	}

	var nombreProducto string

	for _, p := range productos {
		if p.ID == id {
			nombreProducto = p.Nombre
			break
		}
	}

	productos = nuevosProductos

	pedido := models.Pedido{
		ID:             len(pedidos) + 1,
		ClienteID:      cliente.ID,
		ProductoID:     id,
		NombreProducto: nombreProducto,
		Cantidad:       cantidad,
		Total:          total,
	}

	pedidos = services.RegistrarPedido(pedidos, pedido)

	fmt.Println("\nPEDIDO GENERADO CORRECTAMENTE")
	fmt.Println("Cliente:", cliente.Nombre)
	fmt.Println("Producto:", nombreProducto)
	fmt.Println("Cantidad:", cantidad)
	fmt.Printf("Total: $%.2f\n", total)
}

func listarPedidos() {
	fmt.Println("\n--- HISTORIAL DE PEDIDOS ---")

	if len(pedidos) == 0 {
		fmt.Println("No existen pedidos registrados.")
		return
	}

	for _, p := range services.ListarPedidos(pedidos) {
		fmt.Printf("Pedido: %d | Cliente: %d | Producto: %s | Cantidad: %d | Total: $%.2f\n",
			p.ID, p.ClienteID, p.NombreProducto, p.Cantidad, p.Total)
	}
}

func menu() {
	for {
		fmt.Println("\n================================")
		fmt.Println(" SISTEMA DE GESTION E-COMMERCE")
		fmt.Println("================================")
		fmt.Println("1. Registrar producto")
		fmt.Println("2. Listar productos")
		fmt.Println("3. Buscar producto")
		fmt.Println("4. Actualizar inventario")
		fmt.Println("5. Registrar cliente")
		fmt.Println("6. Listar clientes")
		fmt.Println("7. Crear pedido")
		fmt.Println("8. Historial de pedidos")
		fmt.Println("9. Salir")

		opcion := leerTexto("Seleccione una opcion: ")

		switch opcion {
		case "1":
			registrarProducto()
		case "2":
			listarProductos()
		case "3":
			buscarProducto()
		case "4":
			actualizarInventario()
		case "5":
			registrarCliente()
		case "6":
			listarClientes()
		case "7":
			crearPedido()
		case "8":
			listarPedidos()
		case "9":
			fmt.Println("Gracias por utilizar el sistema.")
			return
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}

func main() {
	menu()
}
