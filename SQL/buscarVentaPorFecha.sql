-- Buscar ventas en una fecha especifica
SELECT ID_Producto, ID_Cliente, Cantidad, Fecha FROM Ventas WHERE Fecha = ?;