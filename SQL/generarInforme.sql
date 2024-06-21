-- Generar un informe de productos mas vendidos en un periodo de tiempo especifico
SELECT ID_Producto, SUM(Cantidad) as Total_Vendido FROM Ventas WHERE Fecha BETWEEN ? AND ? GROUP BY ID_Producto ORDER BY Total_Vendido DESC;
