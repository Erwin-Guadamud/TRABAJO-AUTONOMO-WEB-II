package database

import (
    "fmt"
    "os"
    "TRABAJOAUTOWEB/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=prefer",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )

    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    DB = database
    fmt.Println("Database connection successfully opened")
    return nil
}

func MigrateTables() {
    DB.AutoMigrate(&models.Actores{}, &models.ActoresPelicula{}, &models.Comentarios{}, &models.Genero{}, &models.Idioma{}, &models.PeliculaGenero{}, &models.Peliculas{}, &models.Usuario{})
}