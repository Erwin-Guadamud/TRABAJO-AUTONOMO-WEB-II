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

var usuarios []Usuario // Definición del slice para almacenar los usuarios

type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
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
	http.HandleFunc("/usuarios", obtenerUsuarios)
	http.HandleFunc("/usuario", crearUsuario)
	http.HandleFunc("/usuario/path", actualizarUsuario)
	http.HandleFunc("/usuario/delete", eliminarUsuario)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor en ejecución en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// obtenerUsuarios obtiene todos los usuarios
func obtenerUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

// crearUsuario crea un nuevo usuario
func crearUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario Usuario
	_ = json.NewDecoder(r.Body).Decode(&usuario)
	usuarios = append(usuarios, usuario)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuario)
}

// actualizarUsuario actualiza un usuario existente
func actualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Path[len("/usuario/"):]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID de usuario no válido", http.StatusBadRequest)
		return
	}

	var usuarioActualizado Usuario
	_ = json.NewDecoder(r.Body).Decode(&usuarioActualizado)

	for index, usuario := range usuarios {
		if usuario.ID == id {
			usuarios[index] = usuarioActualizado
			json.NewEncoder(w).Encode(usuarioActualizado)
			return
		}
	}

	http.Error(w, "Usuario no encontrado", http.StatusNotFound)
}

// eliminarUsuario elimina un usuario existente
func eliminarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Path[len("/usuario/"):]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID de usuario no válido", http.StatusBadRequest)
		return
	}

	for index, usuario := range usuarios {
		if usuario.ID == id {
			usuarios = append(usuarios[:index], usuarios[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Usuario no encontrado", http.StatusNotFound)
}