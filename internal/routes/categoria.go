package routes

import (
	"encoding/json"
	"myapi/internal/config"
	"myapi/internal/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CategoriaRoutes(r *mux.Router) {
	r.HandleFunc("/categorias", listCategoriasHandler)
	r.HandleFunc("/categorias/get", getCategoriaHandler)
	r.HandleFunc("/categorias/create", createCategoriaHandler)
	r.HandleFunc("/categorias/update", updateCategoriaHandler)
	r.HandleFunc("/categorias/delete", deleteCategoriaHandler)
}

// ==================== HANDLERS PARA CATEGORIAS ====================

// Listar todas as categorias
func listCategoriasHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var categorias []models.Categoria
	if err := config.DB.Find(&categorias).Error; err != nil {
		http.Error(w, "Erro ao buscar categorias", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categorias)
}

// Buscar uma única categoria pelo id (via query string: ?id=1)
func getCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID não fornecido", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	var categoria models.Categoria
	if err := config.DB.First(&categoria, id).Error; err != nil {
		http.Error(w, "Categoria não encontrada", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(categoria)
}

// Criar uma nova categoria (envie JSON via POST)
func createCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var categoria models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&categoria); err != nil {
		http.Error(w, "Erro ao decodificar a categoria", http.StatusBadRequest)
		return
	}
	if err := config.DB.Create(&categoria).Error; err != nil {
		http.Error(w, "Erro ao criar a categoria", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categoria)
}

// Atualizar uma categoria (envie JSON via PUT, com o campo id preenchido)
func updateCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var categoria models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&categoria); err != nil {
		http.Error(w, "Erro ao decodificar a categoria", http.StatusBadRequest)
		return
	}
	if err := config.DB.Save(&categoria).Error; err != nil {
		http.Error(w, "Erro ao atualizar a categoria", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categoria)
}

// Deletar uma categoria (via query string: ?id=1)
func deleteCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID não fornecido", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	if err := config.DB.Delete(&models.Categoria{}, id).Error; err != nil {
		http.Error(w, "Erro ao deletar a categoria", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Categoria deletada com sucesso"))
}
