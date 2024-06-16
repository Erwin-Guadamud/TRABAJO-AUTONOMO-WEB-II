package controllers

import (
	"github.com/Erwin-Guadamud/TRABAJO-AUTONOMO-WEB-II/models"
	"github.com/gofiber/fiber/v2"
)

// Controladores para Producto

// Obtener todos los productos
func GetAllProductos(c *fiber.Ctx) error {
	var productos []models.Producto
	result := models.DB.Find(&productos)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(productos)
}

// Obtener un producto por su ID
func GetProductoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var producto models.Producto
	result := models.DB.First(&producto, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Producto no encontrado"})
	}
	return c.JSON(producto)
}

// Crear un nuevo producto
func CreateProducto(c *fiber.Ctx) error {
	var producto models.Producto
	if err := c.BodyParser(&producto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al decodificar la solicitud"})
	}

	// Verificar si todos los campos del producto están llenos
	if producto.Nombre == "" || producto.Descripcion == "" || producto.PrecioCompra == 0 || producto.PrecioVenta == 0 || producto.Marca == "" || producto.CategoriaID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Todos los campos del producto son obligatorios"})
	}

	result := models.DB.Create(&producto)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(producto)
}

// Actualizar un producto existente
func UpdateProducto(c *fiber.Ctx) error {
	id := c.Params("id")
	var producto models.Producto
	// Buscar el producto por su ID
	result := models.DB.First(&producto, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Producto no encontrado"})
	}

	// Parsear los datos del cuerpo de la solicitud para actualizar el producto
	if err := c.BodyParser(&producto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al decodificar la solicitud"})
	}

	// Guardar los cambios en la base de datos
	models.DB.Save(&producto)
	return c.JSON(producto)
}

// Eliminar un producto existente
func DeleteProducto(c *fiber.Ctx) error {
	id := c.Params("id")
	var producto models.Producto
	result := models.DB.Delete(&producto, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Producto no encontrado"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
