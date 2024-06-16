package controllers

import (
	"github.com/Erwin-Guadamud/TRABAJO-AUTONOMO-WEB-II/models"
	"github.com/gofiber/fiber/v2"
)

// Controladores para Categoría

// Obtener todas las categorías
func GetAllCategories(c *fiber.Ctx) error {
	var categorias []models.Categoria
	result := models.DB.Find(&categorias)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(categorias)
}

// Obtener una categoría por su ID
func GetCategoryByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var categoria models.Categoria
	result := models.DB.First(&categoria, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Categoría no encontrada"})
	}
	return c.JSON(categoria)
}

// Crear una nueva categoría
func CreateCategory(c *fiber.Ctx) error {
	var categoria models.Categoria
	if err := c.BodyParser(&categoria); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al decodificar la solicitud"})
	}

	// Verificar si todos los campos de la categoría están llenos
	if categoria.Nombre == "" || categoria.Descripcion == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Todos los campos de la categoría son obligatorios"})
	}

	result := models.DB.Create(&categoria)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(categoria)
}

// Actualizar una categoría existente
func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var categoria models.Categoria
	// Buscar la categoría por su ID
	result := models.DB.First(&categoria, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Categoría no encontrada"})
	}

	// Parsear los datos del cuerpo de la solicitud para actualizar la categoría
	if err := c.BodyParser(&categoria); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al decodificar la solicitud"})
	}

	// Guardar los cambios en la base de datos
	models.DB.Save(&categoria)
	return c.JSON(categoria)
}

// Eliminar una categoría existente
func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var categoria models.Categoria
	result := models.DB.Delete(&categoria, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Categoría no encontrada"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
