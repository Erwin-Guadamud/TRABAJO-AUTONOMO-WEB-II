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

//Model categoria 


type Categoria struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}


//Model proveedor


type Proveedor struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Direccion string `json:"direccion"`
	Telefono string `json:"telefono"`
	Contacto string `json:"contacto"`
}