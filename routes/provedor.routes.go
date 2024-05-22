import (
    "encoding/json"
    "net/http"
    "strconv"
    "myapp/models"
    "myapp/database"
)

// obtenerProveedores obtiene todos los proveedores
func obtenerProveedores(w http.ResponseWriter, r *http.Request) {
    var proveedores []models.Proveedor
    result := database.DB.Find(&proveedores)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(proveedores)
}

// crearProveedor crea un nuevo proveedor
func crearProveedor(w http.ResponseWriter, r *http.Request) {
    var proveedor models.Proveedor
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

    result := database.DB.Create(&proveedor)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }
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

    var proveedor models.Proveedor
    if err := database.DB.First(&proveedor, id).Error; err != nil {
        http.Error(w, "Proveedor no encontrado", http.StatusNotFound)
        return
    }

    if err := json.NewDecoder(r.Body).Decode(&proveedor); err != nil {
        http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
        return
    }

    if err := database.DB.Save(&proveedor).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(proveedor)
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
