package models

import (
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB // Exportamos la variable DB

// Función para inicializar la conexión con la base de datos
func InitDB(database *gorm.DB) {
	DB = database
}

// Modelo Categoria
type Categoria struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Nombre      string `gorm:"not null" json:"nombre"`
	Descripcion string `json:"descripcion"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
