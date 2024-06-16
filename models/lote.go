package models

import "time"

type Lote struct {
	ID       int       `json:"id"`
	Fecha    time.Time `json:"fecha"`
	IDProducto int       `json:"productoID"`
}