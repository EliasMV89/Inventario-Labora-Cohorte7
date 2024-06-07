-- 5. Actualizar la cantidad disponible de un producto después de una venta
UPDATE Productos SET Cantidad = Cantidad - cantidad_vendida WHERE ID = id_producto;
