package models

import (
    "gorm.io/gorm"
    "time"
)

type Compra struct {
    gorm.Model
    ID        int `gorm:"primaryKey"`
    Fecha     time.Time
    Proveedor string
    Detalles  []DetalleCompra `gorm:"foreignKey:IDCompra"`
}