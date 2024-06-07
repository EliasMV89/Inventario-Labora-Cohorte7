-- 2. Buscar producto por nombre, categoria o proveedor
SELECT * FROM Productos
WHERE Productos.Nombre LIKE '%nombre_producto%'
OR Productos.Categoria LIKE '%nombre_categoria%'
OR Productos.ID_Proveedor = id_proveedor;
