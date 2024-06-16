package controllers

import (
	"time"

	"github.com/Erwin-Guadamud/TRABAJO-AUTONOMO-WEB-II/models"
	"github.com/gofiber/fiber/v2"
)

// Controladores para Compra

// Obtener todas las compras
func GetAllCompras(c *fiber.Ctx) error {
	var compras []models.Compra
	result := models.DB.Preload("Detalles").Find(&compras)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(compras)
}

// Obtener una compra por su ID
func GetCompraByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var compra models.Compra
	result := models.DB.Preload("Detalles").First(&compra, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Compra no encontrada"})
	}
	return c.JSON(compra)
}

// Crear una nueva compra
func CreateCompra(c *fiber.Ctx) error {
	var input struct {
		Fecha       string `json:"Fecha"`
		IDProveedor int    `json:"IDProveedor"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al decodificar la solicitud"})
	}

	fecha, err := time.Parse("2006-01-02", input.Fecha)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Formato de fecha inválido"})
	}

	compra := models.NewCompra(fecha, input.IDProveedor)

	result := models.DB.Create(&compra)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(compra)
}

// Actualizar una compra existente
func UpdateCompra(c *fiber.Ctx) error {
	id := c.Params("id")
	var compra models.Compra
	// Buscar la compra por su ID
	result := models.DB.Preload("Detalles").First(&compra, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Compra no encontrada"})
	}

	// Parsear los datos del cuerpo de la solicitud para actualizar la compra
	if err := c.BodyParser(&compra); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al decodificar la solicitud"})
	}

	// Guardar los cambios en la base de datos
	models.DB.Save(&compra)
	return c.JSON(compra)
}

// Eliminar una compra existente
func DeleteCompra(c *fiber.Ctx) error {
	id := c.Params("id")
	var compra models.Compra
	result := models.DB.Delete(&compra, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Compra no encontrada"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
