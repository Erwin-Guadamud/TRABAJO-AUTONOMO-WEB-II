package models

import (
	"time"

	"gorm.io/gorm"
)

type Compra struct {
	gorm.Model
	Fecha       time.Time `json:"Fecha"`
	IDProveedor int       `json:"IDProveedor"`
	Proveedor   Proveedor `gorm:"foreignKey:IDProveedor"` // Relación con la tabla Proveedor
}

func NewCompra(fecha time.Time, idProveedor int) *Compra {
	return &Compra{
		Fecha:       fecha,
		IDProveedor: idProveedor,
	}
}
