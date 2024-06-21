-- Buscar producto por nombre o categoria
SELECT Id, Nombre, Categoria, Precio, Cantidad, ID_Proveedor FROM Productos WHERE Nombre LIKE ? OR Categoria LIKE ?;

