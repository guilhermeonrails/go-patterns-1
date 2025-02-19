package main

import (
	"encoding/json"
	"fmt"
	"log"
	"myapi/internal/config"
	"myapi/internal/handlers"
	"myapi/internal/models"
	"net/http"
	"strconv"
)

func main() {

	config.ConnectDatabase()

	// Endpoint raiz
	http.HandleFunc("/api", indexHandler)

	// Endpoints para Itens
	http.HandleFunc("/itens", handlers.ListItensHandler)                // GET para listar todos os itens
	http.HandleFunc("/itens/get", handlers.GetItenHandler)              // GET para buscar um item (espera id via query: ?id=1)
	http.HandleFunc("/itens/get-code", handlers.GetItenByCodigoHandler) // get-code?codigo=TEC001
	http.HandleFunc("/itens/create", handlers.CreateItenHandler)        // POST para criar um item
	http.HandleFunc("/itens/update", handlers.UpdateItenHandler)        // PUT para atualizar um item (JSON com id)
	http.HandleFunc("/itens/delete", handlers.DeleteItenHandler)        // DELETE para deletar um item (espera id via query: ?id=1)

	// Endpoints para Categorias
	http.HandleFunc("/categorias", listCategoriasHandler)         // GET para listar todas as categorias
	http.HandleFunc("/categorias/get", getCategoriaHandler)       // GET para buscar uma categoria (espera id via query)
	http.HandleFunc("/categorias/create", createCategoriaHandler) // POST para criar uma categoria
	http.HandleFunc("/categorias/update", updateCategoriaHandler) // PUT para atualizar uma categoria (JSON com id)
	http.HandleFunc("/categorias/delete", deleteCategoriaHandler) // DELETE para deletar uma categoria (espera id via query)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler raiz
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API Go!")
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
