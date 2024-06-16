package models

import (
	"gorm.io/gorm"
)

type Proveedor struct {
	gorm.Model
	ID        int    `gorm:"primaryKey"`
	Nombre    string
	Direccion string
	Telefono  string
	Contacto  string
}