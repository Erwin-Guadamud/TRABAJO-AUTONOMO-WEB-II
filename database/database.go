package database

import (
	"fmt"

	"github.com/Erwin-Guadamud/TRABAJO-AUTONOMO-WEB-II/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=prefer"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
	fmt.Println("Database connection successfully opened")
	MigrateTables()

}
func MigrateTables() {
	DB.AutoMigrate(&models.Categoria{}, &models.Compra{}, &models.DetalleCompra{}, &models.Lote{}, &models.Producto{}, &models.Proveedor{})
}
