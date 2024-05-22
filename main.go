package main

import (
    "log"
    "trabajoautoweb/database"
    "trabajoautoweb/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Conectarse a la base de datos
    err := database.ConnectDatabase()
    if err != nil {
        log.Fatalf("Error al conectar a la base de datos: %v", err)
    }

    app := fiber.New()

    // Configurar las rutas
    routes.SetupRoutes(app)

    // Iniciar la aplicación
    err = app.Listen(":3000")
    if err != nil {
        log.Fatalf("Error al iniciar la aplicación: %v", err)
    }
}