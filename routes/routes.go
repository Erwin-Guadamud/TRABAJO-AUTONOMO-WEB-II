package routes

import (
	"github.com/gin-gonic/gin"
    "trabajoautoweb/models"
)

func SetupRoutes() {
	router := gin.Default()

	// Rutas para el modelo de categoría
	categoryRoutes := router.Group("/categories")
	{
		categoryRoutes.GET("/", controllers.GetAllCategories)
		categoryRoutes.GET("/:id", controllers.GetCategoryByID)
		categoryRoutes.POST("/", controllers.CreateCategory)
		categoryRoutes.PUT("/:id", controllers.UpdateCategory)
		categoryRoutes.DELETE("/:id", controllers.DeleteCategory)
	}

	// Agrega aquí las rutas para otros modelos
	//Rutas de productos
	productRoutes := router.Group("/products")
	{
		productRoutes.GET("/", controllers.GetAllProducts)
		productRoutes.GET("/:id", controllers.GetProductByID)
		productRoutes.POST("/", controllers.CreateProduct)
		productRoutes.PUT("/:id", controllers.UpdateProduct)
		productRoutes.DELETE("/:id", controllers.DeleteProduct)
	}

	router.Run(":8080")
}