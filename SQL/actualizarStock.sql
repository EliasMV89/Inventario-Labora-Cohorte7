-- Actualizar la cantidad disponible de un producto especifico despues de una venta
UPDATE Productos SET Cantidad = Cantidad - cantidad_vendida WHERE ID = id_producto;
