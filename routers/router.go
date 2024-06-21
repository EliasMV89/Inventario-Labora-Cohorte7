/*package routers

import (
	"Inventario/controllers"
	"Inventario/utils"
	

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Clientes
	router.HandleFunc("/clientes", controllers.AgregarCliente).Methods("POST")
	router.HandleFunc("/clientes", controllers.ListarClientes).Methods("GET")
	router.HandleFunc("/cliente/{id}", controllers.ModificarCliente).Methods("PUT")
	router.HandleFunc("/cliente/{id}", controllers.EliminarCliente).Methods("DELETE")

	// Productos
	router.HandleFunc("/productos", controllers.AgregarProducto).Methods("POST")
	router.HandleFunc("/producto/{id}", controllers.ModificarProducto).Methods("PUT")
	router.HandleFunc("/producto/{id}", controllers.EliminarProducto).Methods("DELETE")
	router.HandleFunc("/productos", controllers.ListarProductos).Methods("GET")
	router.HandleFunc("/productos/buscar", controllers.BuscarProducto).Methods("GET")

	// Proveedores
	router.HandleFunc("/proveedores", controllers.AgregarProveedor).Methods("POST")
	router.HandleFunc("/proveedores", controllers.ListarProveedores).Methods("GET")
	router.HandleFunc("/proveedor/{id}", controllers.ModificarProveedor).Methods("PUT")
	router.HandleFunc("/proveedor/{id}", controllers.EliminarProveedor).Methods("DELETE")

	// Ventas
	router.HandleFunc("/ventas", controllers.AgregarVenta).Methods("POST")
	router.HandleFunc("/ventas", controllers.ListarVentas).Methods("GET")
	router.HandleFunc("/ventas/buscar", controllers.BuscarVentaPorFecha).Methods("GET")
	router.HandleFunc("/ventas/informe", controllers.GenerarInforme).Methods("GET")

	router.Use(utils.CorsMiddleware)

	return router
}*/

package routers

import (
    "Inventario/controllers"
    "Inventario/utils"

    "github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)

    // Clientes
    router.HandleFunc("/clientes", controllers.AgregarCliente).Methods("POST", "OPTIONS")
    router.HandleFunc("/clientes", controllers.ListarClientes).Methods("GET", "OPTIONS")
    router.HandleFunc("/cliente/{id}", controllers.ModificarCliente).Methods("PUT", "OPTIONS")
    router.HandleFunc("/cliente/{id}", controllers.EliminarCliente).Methods("DELETE", "OPTIONS")

    // Productos
    router.HandleFunc("/productos", controllers.AgregarProducto).Methods("POST", "OPTIONS")
    router.HandleFunc("/producto/{id}", controllers.ModificarProducto).Methods("PUT", "OPTIONS")
    router.HandleFunc("/producto/{id}", controllers.EliminarProducto).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/productos", controllers.ListarProductos).Methods("GET", "OPTIONS")
    router.HandleFunc("/productos/buscar", controllers.BuscarProducto).Methods("GET", "OPTIONS")

    // Proveedores
    router.HandleFunc("/proveedores", controllers.AgregarProveedor).Methods("POST", "OPTIONS")
    router.HandleFunc("/proveedores", controllers.ListarProveedores).Methods("GET", "OPTIONS")
    router.HandleFunc("/proveedor/{id}", controllers.ModificarProveedor).Methods("PUT", "OPTIONS")
    router.HandleFunc("/proveedor/{id}", controllers.EliminarProveedor).Methods("DELETE", "OPTIONS")

    // Ventas
    router.HandleFunc("/ventas", controllers.AgregarVenta).Methods("POST", "OPTIONS")
    router.HandleFunc("/ventas", controllers.ListarVentas).Methods("GET", "OPTIONS")
    router.HandleFunc("/ventas/buscar", controllers.BuscarVentaPorFecha).Methods("GET", "OPTIONS")
    router.HandleFunc("/ventas/informe", controllers.GenerarInforme).Methods("GET", "OPTIONS")

    // Middleware CORS
    router.Use(utils.CorsMiddleware)

    return router
}
