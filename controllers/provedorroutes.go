package controllers

import (
	"github.com/Erwin-Guadamud/TRABAJO-AUTONOMO-WEB-II/models"
	"github.com/gofiber/fiber/v2"
)

// Controladores para Proveedor

// Obtener todos los proveedores
func GetAllProveedores(c *fiber.Ctx) error {
	var proveedores []models.Proveedor
	result := models.DB.Find(&proveedores)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(proveedores)
}

// Obtener un proveedor por su ID
func GetProveedorByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var proveedor models.Proveedor
	result := models.DB.First(&proveedor, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Proveedor no encontrado"})
	}
	return c.JSON(proveedor)
}

// Crear un nuevo proveedor
func CreateProveedor(c *fiber.Ctx) error {
	var proveedor models.Proveedor
	if err := c.BodyParser(&proveedor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al decodificar la solicitud"})
	}

	// Verificar si todos los campos del proveedor están llenos
	if proveedor.Nombre == "" || proveedor.Direccion == "" || proveedor.Telefono == "" || proveedor.Contacto == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Todos los campos del proveedor son obligatorios"})
	}

	result := models.DB.Create(&proveedor)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(proveedor)
}

// Actualizar un proveedor existente
func UpdateProveedor(c *fiber.Ctx) error {
	id := c.Params("id")
	var proveedor models.Proveedor
	// Buscar el proveedor por su ID
	result := models.DB.First(&proveedor, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Proveedor no encontrado"})
	}

	// Parsear los datos del cuerpo de la solicitud para actualizar el proveedor
	if err := c.BodyParser(&proveedor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al decodificar la solicitud"})
	}

	// Guardar los cambios en la base de datos
	models.DB.Save(&proveedor)
	return c.JSON(proveedor)
}

// Eliminar un proveedor existente
func DeleteProveedor(c *fiber.Ctx) error {
	id := c.Params("id")
	var proveedor models.Proveedor
	result := models.DB.Delete(&proveedor, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Proveedor no encontrado"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
