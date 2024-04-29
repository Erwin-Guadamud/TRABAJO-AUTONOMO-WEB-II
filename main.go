package main

import (
	// "database/sql"
	// "fmt"
	// "log"
	// _ "github.com/denisenkom/go-mssqldb"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)



//Model categoria 

var categorias []Categoria // Slice para almacenar las categorías

type Categoria struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}


//Model proveedor
var proveedores []Proveedor // Definición del slice para almacenar los proveedores

type Proveedor struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Direccion string `json:"direccion"`
	Telefono string `json:"telefono"`
	Contacto string `json:"contacto"`
}

func main() {
	///Conexion a SQL

	// server := "SQLEXPRESS"
	// port := 1433
	// user := "pc"
	// password := "010304"
	// database := "GestionCompra"

	// // Cadena de conexión
	// connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", server, user, password, port, database)

	// // Conectarse a la base de datos
	// db, err := sql.Open("sqlserver", connectionString)
	// if err != nil {
	// 	log.Fatal("Error al conectar a la base de datos:", err.Error())
	// }
	// defer db.Close()

	// // Verificar la conexión
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal("Error al verificar la conexión:", err.Error())
	// }

	// fmt.Println("Conexión exitosa a la base de datos SQL Server")

	
	// Rutas proveedor CRUD API-REST
	http.HandleFunc("/proveedores", obtenerProveedores)
	http.HandleFunc("/proveedor/create", crearProveedor)
	http.HandleFunc("/proveedor/put/", actualizarProveedor)
	http.HandleFunc("/proveedor/delete/", eliminarProveedor)

	// Rutas Categoria CRUD API-REST
	http.HandleFunc("/categorias", obtenerCategorias)
	http.HandleFunc("/categoria/create", crearCategoria)
	http.HandleFunc("/categoria/put/", actualizarCategoria)
	http.HandleFunc("/categoria/delete/", eliminarCategoria)
		
	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor en ejecución en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


/////
/// CRUD API-REST PROVEEDOR
//

// obtenerProveedores obtiene todos los proveedores
func obtenerProveedores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proveedores)
}

// crearProveedor crea un nuevo proveedor
func crearProveedor(w http.ResponseWriter, r *http.Request) {
	var proveedor Proveedor
	err := json.NewDecoder(r.Body).Decode(&proveedor)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	// Verificar si todos los campos del proveedor están llenos
	if proveedor.Nombre == "" || proveedor.Direccion == "" || proveedor.Telefono == "" || proveedor.Contacto == "" {
		http.Error(w, "Todos los campos del proveedor son obligatorios", http.StatusBadRequest)
		return
	}

	proveedores = append(proveedores, proveedor)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proveedor)
}


// actualizarProveedor actualiza un proveedor existente
func actualizarProveedor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Path[len("/proveedor/"):]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID de proveedor no válido", http.StatusBadRequest)
		return
	}

	var proveedorActualizado Proveedor
	_ = json.NewDecoder(r.Body).Decode(&proveedorActualizado)

	for index, proveedor := range proveedores {
		if proveedor.ID == id {
			proveedores[index] = proveedorActualizado
			json.NewEncoder(w).Encode(proveedorActualizado)
			return
		}
	}

	http.Error(w, "Proveedor no encontrado", http.StatusNotFound)
}

// eliminarProveedor elimina un proveedor existente
func eliminarProveedor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Path[len("/proveedor/"):]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID de proveedor no válido", http.StatusBadRequest)
		return
	}

	for index, proveedor := range proveedores {
		if proveedor.ID == id {
			proveedores = append(proveedores[:index], proveedores[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Proveedor no encontrado", http.StatusNotFound)
}


/////
/// CRUD API-REST Categoria
//

// obtenerCategorias obtiene todas las categorías
func obtenerCategorias(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categorias)
}

// crearCategoria crea una nueva categoría
func crearCategoria(w http.ResponseWriter, r *http.Request) {
	var categoria Categoria
	err := json.NewDecoder(r.Body).Decode(&categoria)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	// Verificar si todos los campos de la categoría están llenos
	if categoria.Nombre == "" || categoria.Descripcion == "" {
		http.Error(w, "Todos los campos de la categoría son obligatorios", http.StatusBadRequest)
		return
	}

	// Generar ID para la nueva categoría
	if len(categorias) == 0 {
		categoria.ID = 1
	} else {
		lastCategoria := categorias[len(categorias)-1]
		categoria.ID = lastCategoria.ID + 1
	}

	categorias = append(categorias, categoria)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categoria)
}

// actualizarCategoria actualiza una categoría existente
func actualizarCategoria(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Path[len("/categoria/"):]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID de categoría no válido", http.StatusBadRequest)
		return
	}

	var categoriaActualizada Categoria
	err = json.NewDecoder(r.Body).Decode(&categoriaActualizada)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	for index, categoria := range categorias {
		if categoria.ID == id {
			categorias[index] = categoriaActualizada
			json.NewEncoder(w).Encode(categoriaActualizada)
			return
		}
	}

	http.Error(w, "Categoría no encontrada", http.StatusNotFound)
}

// eliminarCategoria elimina una categoría existente
func eliminarCategoria(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Path[len("/categoria/"):]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID de categoría no válido", http.StatusBadRequest)
		return
	}

	for index, categoria := range categorias {
		if categoria.ID == id {
			categorias = append(categorias[:index], categorias[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Categoría no encontrada", http.StatusNotFound)
}


