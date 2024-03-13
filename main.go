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

var proveedores []Proveedor // Definición del slice para almacenar los proveedores

type Proveedor struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Direccion string `json:"direccion"`
	Telefono string `json:"telefono"`
	Contacto string `json:"contacto"`
}

func main() {
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

	
	// Rutas
	http.HandleFunc("/proveedores", obtenerProveedores)
	http.HandleFunc("/proveedor/created", crearProveedor)
	http.HandleFunc("/proveedor/put", actualizarProveedor)
	http.HandleFunc("/proveedor/delete", eliminarProveedor)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor en ejecución en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

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