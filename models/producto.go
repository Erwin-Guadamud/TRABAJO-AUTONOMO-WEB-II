package models

import (
    "gorm.io/gorm"
)

type Producto struct {
    gorm.Model
    ID            int     `gorm:"primaryKey"`
    Nombre        string
    Descripcion   string
    PrecioCompra  float64
    PrecioVenta   float64
    Marca         string
    CategoriaID   int
    Categoria     Categoria `gorm:"foreignKey:CategoriaID"`
}

func NewProducto(id int, nombre string, descripcion string, precioCompra float64, precioVenta float64, marca string, categoriaID int) *Producto {
    return &Producto{
        ID:            id,
        Nombre:        nombre,
        Descripcion:   descripcion,
        PrecioCompra:  precioCompra,
        PrecioVenta:   precioVenta,
        Marca:         marca,
        CategoriaID:   categoriaID,
    }
}