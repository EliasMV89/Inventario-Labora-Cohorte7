-- 6. Generar un informe de productos más vendidos en un período de tiempo específico
SELECT ID_Producto, SUM(Cantidad) as Total_Vendido
FROM Ventas
WHERE Fecha BETWEEN 'fecha_inicio' AND 'fecha_fin'
GROUP BY ID_Producto
ORDER BY Total_Vendido DESC;
