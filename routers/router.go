package routers

import (
    "Inventario/controllers" 
    "Inventario/utils"       

    "github.com/gorilla/mux" 
)

// InitRouter inicializa el enrutador y configura las rutas de la aplicacion
func InitRouter() *mux.Router {
    // Crea un nuevo enrutador y establece StrictSlash en true para manejar correctamente las URL con o sin barra al final
    router := mux.NewRouter().StrictSlash(true)

    // Rutas para Clientes
    router.HandleFunc("/clientes", controllers.AgregarCliente).Methods("POST", "OPTIONS")      // Ruta para agregar un cliente
    router.HandleFunc("/clientes", controllers.ListarClientes).Methods("GET", "OPTIONS")       // Ruta para listar todos los clientes
    router.HandleFunc("/cliente/{id}", controllers.ModificarCliente).Methods("PUT", "OPTIONS") // Ruta para modificar un cliente por ID
    router.HandleFunc("/cliente/{id}", controllers.EliminarCliente).Methods("DELETE", "OPTIONS") // Ruta para eliminar un cliente por ID

    // Rutas para Productos
    router.HandleFunc("/productos", controllers.AgregarProducto).Methods("POST", "OPTIONS")      // Ruta para agregar un producto
    router.HandleFunc("/producto/{id}", controllers.ModificarProducto).Methods("PUT", "OPTIONS") // Ruta para modificar un producto por ID
    router.HandleFunc("/producto/{id}", controllers.EliminarProducto).Methods("DELETE", "OPTIONS") // Ruta para eliminar un producto por ID
    router.HandleFunc("/productos", controllers.ListarProductos).Methods("GET", "OPTIONS")       // Ruta para listar todos los productos
    router.HandleFunc("/productos/buscar", controllers.BuscarProducto).Methods("GET", "OPTIONS") // Ruta para buscar productos

    // Rutas para Proveedores
    router.HandleFunc("/proveedores", controllers.AgregarProveedor).Methods("POST", "OPTIONS")      // Ruta para agregar un proveedor
    router.HandleFunc("/proveedores", controllers.ListarProveedores).Methods("GET", "OPTIONS")       // Ruta para listar todos los proveedores
    router.HandleFunc("/proveedor/{id}", controllers.ModificarProveedor).Methods("PUT", "OPTIONS") // Ruta para modificar un proveedor por ID
    router.HandleFunc("/proveedor/{id}", controllers.EliminarProveedor).Methods("DELETE", "OPTIONS") // Ruta para eliminar un proveedor por ID

    // Rutas para Ventas
    router.HandleFunc("/ventas", controllers.AgregarVenta).Methods("POST", "OPTIONS")              // Ruta para agregar una venta
    router.HandleFunc("/ventas", controllers.ListarVentas).Methods("GET", "OPTIONS")               // Ruta para listar todas las ventas
    router.HandleFunc("/ventas/buscar", controllers.BuscarVentaPorFecha).Methods("GET", "OPTIONS") // Ruta para buscar ventas por fecha
    router.HandleFunc("/ventas/informe", controllers.GenerarInforme).Methods("GET", "OPTIONS")     // Ruta para generar un informe de ventas

    // Middleware CORS
    router.Use(utils.CorsMiddleware)

    return router
}

    