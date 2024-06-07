package main

import (
	"Inventario/servicios"
	"Inventario/utils"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// Establece la conexión con la base de datos
	db, err := utils.ConectarBaseDeDatos()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Opciones del sistema
	for {
		fmt.Print("Bienvenido!")
		fmt.Print("*******************************************\n")
		fmt.Print("Elige una opcion\n")

		fmt.Print("*******************************************\n")
		fmt.Println("1. Agregar un producto")
		fmt.Println("2. Agregar un proveedor")
		fmt.Println("3. Agregar un cliente")
		fmt.Println("4. Listar todos los productos")
		fmt.Println("5. Buscar productos por nombre, categoría o proveedor")
		fmt.Println("6. Registrar una nueva venta de un producto a un cliente")
		fmt.Println("7. Mostrar las ventas realizadas en un día específico")
		fmt.Println("8. Generar un informe de productos más vendidos en un período de tiempo específico")
		fmt.Println("9. Salir del sistema")
		fmt.Println("*********************************")
		fmt.Print("Ingrese su opcion: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Agregar un producto")
			fmt.Printf("Ingrese el nombre del producto: ")
			var nombre string
			fmt.Scanln(&nombre)
			fmt.Printf("Ingrese la categoria del producto: ")
			var categoria string
			fmt.Scanln(&categoria)
			var cantidad int
			for {
				fmt.Printf("Ingrese la cantidad del producto mayor que 0: ")
				fmt.Scanln(&cantidad)
				if cantidad > 0 {
					break
				}
			}
			fmt.Printf("Ingrese el ID del proveedor: ")
			var idProveedor int
			fmt.Scanln(&idProveedor)
			nuevoProducto := servicios.Producto{
				Nombre:       nombre,
				Categoria:    categoria,
				Cantidad:     cantidad,
				ID_Proveedor: idProveedor,
			}
			servicios.AgregarProducto(db, nuevoProducto)
		case 2:
			fmt.Println("Agregar un proveedor")
			fmt.Printf("Ingrese nombre del proveedor: ")
			var nombre string
			fmt.Scanln(&nombre)
			fmt.Printf("Ingrese el telefono del proveedor: ")
			var contacto string
			fmt.Scanln(&contacto)
			nuevoProveedor := servicios.Proveedor{
				Nombre:   nombre,
				Contacto: contacto,
			}
			servicios.AgregarProveedor(db, nuevoProveedor)
		case 3:
			fmt.Println("Agregar un nuevo cliente")
			fmt.Printf("Ingrese nombre del cliente: ")
			var nombre string
			fmt.Scanln(&nombre)
			fmt.Printf("Ingrese el telefono del cliente: ")
			var contacto string
			fmt.Scanln(&contacto)
			nuevoCliente := servicios.Cliente{
				Nombre:   nombre,
				Contacto: contacto,
			}
			servicios.AgregarCliente(db, nuevoCliente)
		case 4:
			fmt.Println("Listar todos los productos")
			servicios.ListarProductos(db)
		case 5:
			fmt.Println("Buscar productos por nombre, categoría o proveedor")
			fmt.Printf("Ingrese el nombre, categoria o ID del proveedor que desea buscar: ")
			var buscar string
			fmt.Scanln(&buscar)
			servicios.BuscarProducto(db, buscar)
		case 6:
			fmt.Println("Registrar una nueva venta de un producto a un cliente")
			fmt.Printf("Ingrese el ID del producto: ")
			var idProducto int
			fmt.Scanln(&idProducto)
			fmt.Printf("Ingrese el ID del cliente: ")
			var idCliente int
			fmt.Scanln(&idCliente)
			fechaActual := time.Now()
			fecha := fechaActual.Format("2006-01-02")
			var cantidad int
			for {
				cantidaStock, err := servicios.ObtenerCantidad(db, idProducto)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Ingrese la cantidad del producto, no mayor a %d : ", cantidaStock)
				fmt.Scanln(&cantidad)
				if cantidaStock < cantidad {
					fmt.Printf("Ingrese la cantidad del producto, no mayor a %d : ", cantidaStock)
					fmt.Scanln(&cantidad)
					continue
				}
				break
			}
			nuevaVenta := servicios.Venta{
				ID_Producto: idProducto,
				ID_Cliente:  idCliente,
				Cantidad:    cantidad,
				Fecha:       fecha,
			}
			servicios.AgregarVenta(db, nuevaVenta)
			servicios.ActualizarStock(db, cantidad, idProducto)
		case 7:
			fmt.Println("Mostrar las ventas realizadas en un día específico")
			fmt.Printf("Ingrese la fecha de busqueda formato(YYYY-MM-DD): ")
			var fecha string
			fmt.Scanln(&fecha)
			servicios.BuscarVentaPorFecha(db, fecha)
		case 8:
			fmt.Println("Generar un informe de productos más vendidos en un período de tiempo específico")
			fmt.Printf("Ingrese la fecha de inicio de la busqueda formato(YYYY-MM-DD): ")
			var fechaInicio string
			fmt.Scanln(&fechaInicio)
			fmt.Printf("Ingrese la fecha de fin de la busqueda formato(YYYY-MM-DD): ")
			var fechaFin string
			fmt.Scanln(&fechaFin)
			servicios.GenerarInforme(db, fechaInicio, fechaFin)
		case 9:
			os.Exit(0)
		default:
			fmt.Println("Opcion invalida")
		}
	}
}
