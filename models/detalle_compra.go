package models

import (
	"gorm.io/gorm"
)

type DetalleCompra struct {
	gorm.Model
	IDCompraDetalle int     `gorm:"primaryKey"`
	Cantidad        int
	PrecioCompra    float64
	IdCompra        int       // Campo de clave foránea
	Compra          Compra    `gorm:"foreignKey:IdCompra"` // Relación con la tabla Compra
}
