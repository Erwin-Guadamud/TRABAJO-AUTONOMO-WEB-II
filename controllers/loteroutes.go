package controllers

import (
	"github.com/Erwin-Guadamud/TRABAJO-AUTONOMO-WEB-II/models"
	"github.com/gofiber/fiber/v2"
)

// Controladores para Lote

// Obtener todos los lotes
func GetAllLotes(c *fiber.Ctx) error {
	var lotes []models.Lote
	result := models.DB.Find(&lotes)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(lotes)
}

// Obtener un lote por su ID
func GetLoteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var lote models.Lote
	result := models.DB.First(&lote, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Lote no encontrado"})
	}
	return c.JSON(lote)
}

// Crear un nuevo lote
func CreateLote(c *fiber.Ctx) error {
	var lote models.Lote
	if err := c.BodyParser(&lote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al decodificar la solicitud"})
	}

	result := models.DB.Create(&lote)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(lote)
}

// Actualizar un lote existente
func UpdateLote(c *fiber.Ctx) error {
	id := c.Params("id")
	var lote models.Lote
	result := models.DB.First(&lote, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Lote no encontrado"})
	}

	if err := c.BodyParser(&lote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al decodificar la solicitud"})
	}

	models.DB.Save(&lote)
	return c.JSON(lote)
}

// Eliminar un lote existente
func DeleteLote(c *fiber.Ctx) error {
	id := c.Params("id")
	var lote models.Lote
	result := models.DB.Delete(&lote, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Lote no encontrado"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
