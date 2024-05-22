package models

import (
    "gorm.io/gorm"
)

type DetalleCompra struct {
    gorm.Model
    IDCompra     int     `gorm:"primaryKey"`
    IDProducto   int     `gorm:"primaryKey"`
    Cantidad     int
    PrecioCompra float64
    Compra       Compra  `gorm:"foreignKey:IDCompra"`
    Producto     Producto `gorm:"foreignKey:IDProducto"`
}