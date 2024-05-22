import (
    "encoding/json"
    "net/http"
    "strconv"
    "myapp/models"
    "myapp/database"
)

// obtenerCategorias obtiene todas las categorías
func obtenerCategorias(w http.ResponseWriter, r *http.Request) {
    var categorias []models.Categoria
    result := database.DB.Find(&categorias)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(categorias)
}

// crearCategoria crea una nueva categoría
func crearCategoria(w http.ResponseWriter, r *http.Request) {
    var categoria models.Categoria
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

    result := database.DB.Create(&categoria)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }
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


