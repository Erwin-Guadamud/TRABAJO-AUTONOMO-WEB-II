package main

import (
	"github.com/Erwin-Guadamud/TRABAJO-AUTONOMO-WEB-II/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Erwin-Guadamud/TRABAJO-AUTONOMO-WEB-II/database"
	"github.com/Erwin-Guadamud/TRABAJO-AUTONOMO-WEB-II/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDatabase()

	// Establece la conexión a la base de datos
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=prefer"), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}
	// Inicializa la conexión a la base de datos en el paquete models
	models.InitDB(db)

	app := fiber.New()

	// Configurar las rutas
	routes.SetupRoutes(app)

	// Iniciar la aplicación
	app.Listen(":3000")
}
