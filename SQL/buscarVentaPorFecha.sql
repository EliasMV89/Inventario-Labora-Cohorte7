-- 4. Mostrar las ventas realizadas en un dia especifo
SELECT ID_Producto, ID_Cliente, Cantidad, Fecha FROM Ventas WHERE Fecha = ?;