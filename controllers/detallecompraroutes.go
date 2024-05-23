package controllers

import (
	"github.com/Erwin-Guadamud/TRABAJO-AUTONOMO-WEB-II/models"
	"github.com/gofiber/fiber/v2"
)

// Controladores para DetalleCompra

// Obtener todos los detalles de compra
func GetAllDetallesCompra(c *fiber.Ctx) error {
	var detalles []models.DetalleCompra
	result := models.DB.Find(&detalles)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(detalles)
}

// Obtener detalles de compra por ID de compra
func GetDetallesCompraByCompraID(c *fiber.Ctx) error {
	id := c.Params("id")
	var detalles []models.DetalleCompra
	result := models.DB.Where("id_compra = ?", id).Find(&detalles)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(detalles)
}

// Crear un nuevo detalle de compra
func CreateDetalleCompra(c *fiber.Ctx) error {
	var detalle models.DetalleCompra
	if err := c.BodyParser(&detalle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al decodificar la solicitud"})
	}

	result := models.DB.Create(&detalle)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(detalle)
}

// Actualizar un detalle de compra existente
func UpdateDetalleCompra(c *fiber.Ctx) error {
	id := c.Params("id")
	var detalle models.DetalleCompra
	result := models.DB.First(&detalle, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Detalle de compra no encontrado"})
	}

	if err := c.BodyParser(&detalle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al decodificar la solicitud"})
	}

	models.DB.Save(&detalle)
	return c.JSON(detalle)
}

// Eliminar un detalle de compra existente
func DeleteDetalleCompra(c *fiber.Ctx) error {
	id := c.Params("id") // Obtener el ID del parámetro de la solicitud
	var detalle models.DetalleCompra
	result := models.DB.Delete(&detalle, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Detalle de compra no encontrado"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
