package routes

import (
	"github.com/Erwin-Guadamud/TRABAJO-AUTONOMO-WEB-II/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Rutas para Categoría
	app.Get("/categories", controllers.GetAllCategories)
	app.Get("/categories/:id", controllers.GetCategoryByID)
	app.Post("/categories", controllers.CreateCategory)
	app.Put("/categories/:id", controllers.UpdateCategory)
	app.Delete("/categories/:id", controllers.DeleteCategory)

	// Rutas para Proovedores
	app.Get("/provedor", controllers.GetAllProveedores)
	app.Get("/provedor/:id", controllers.GetProveedorByID)
	app.Post("/provedor", controllers.CreateProveedor)
	app.Put("/provedor/:id", controllers.UpdateProveedor)
	app.Delete("/provedor/:id", controllers.DeleteProveedor)

	// Rutas para Productos
	app.Get("/productos", controllers.GetAllProductos)
	app.Get("/productos/:id", controllers.GetProductoByID)
	app.Post("/productos", controllers.CreateProducto)
	app.Put("/productos/:id", controllers.UpdateProducto)
	app.Delete("/productos/:id", controllers.DeleteProducto)



	// Rutas para Compras
	app.Get("/compras", controllers.GetAllCompras)
	app.Get("/compras/:id", controllers.GetCompraByID)
	app.Post("/compras", controllers.CreateCompra)
	app.Put("/compras/:id", controllers.UpdateCompra)
	app.Delete("/compras/:id", controllers.DeleteCompra)

	// Rutas para Detalles de Compra
	app.Get("/detalles_compra", controllers.GetAllDetallesCompra)
	app.Get("/detalles_compra/compra/:id", controllers.GetDetallesCompraByCompraID)
	app.Post("/detalles_compra", controllers.CreateDetalleCompra)
	app.Put("/detalles_compra/:id", controllers.UpdateDetalleCompra)
	app.Delete("/detalles_compra/:id", controllers.DeleteDetalleCompra)

		// Rutas para Lotes
	app.Get("/lotes", controllers.GetAllLotes)
	app.Get("/lotes/:id", controllers.GetLoteByID)
	app.Post("/lotes", controllers.CreateLote)
	app.Put("/lotes/:id", controllers.UpdateLote)
	app.Delete("/lotes/:id", controllers.DeleteLote)

}
