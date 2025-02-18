package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Modelo para a tabela "itens"
type Iten struct {
	Id         uint    `gorm:"primaryKey" json:"id"`
	Nome       string  `json:"nome"`
	Codigo     string  `gorm:"unique" json:"codigo"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
}

// Modelo para a tabela "categorias"
type Cat struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	Nome      string `json:"nome"`
	Codigo    string `gorm:"unique" json:"codigo"`
	Descricao string `json:"descricao"`
}

var bd *gorm.DB

func main() {
	// Conexão com o Postgres (usando host "db" pois o docker-compose cria essa rede)
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar com o BD: %v", err)
	}
	bd = db

	// AutoMigrate para criar/ajustar tabelas
	bd.AutoMigrate(&Iten{})
	bd.AutoMigrate(&Cat{})

	// Endpoint raiz
	http.HandleFunc("/api", indexHandler)

	// Endpoints para Itens
	http.HandleFunc("/itens", listItensHandler)                // GET para listar todos os itens
	http.HandleFunc("/itens/get", getItenHandler)              // GET para buscar um item (espera id via query: ?id=1)
	http.HandleFunc("/itens/get-code", getItenByCodigoHandler) // get-code?codigo=TEC001
	http.HandleFunc("/itens/create", createItenHandler)        // POST para criar um item
	http.HandleFunc("/itens/update", updateItenHandler)        // PUT para atualizar um item (JSON com id)
	http.HandleFunc("/itens/delete", deleteItenHandler)        // DELETE para deletar um item (espera id via query: ?id=1)

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

// ==================== HANDLERS PARA ITENS ====================

// Listar todos os itens
func listItensHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var itens []Iten
	if err := bd.Find(&itens).Error; err != nil {
		http.Error(w, "Erro ao buscar itens", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(itens)
}

// Buscar um único item pelo id (via query string: ?id=1)
func getItenHandler(w http.ResponseWriter, r *http.Request) {
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
	var item Iten
	if err := bd.First(&item, id).Error; err != nil {
		http.Error(w, "Item não encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

// Buscar um item pelo campo "codigo"
func getItenByCodigoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	cod := r.URL.Query().Get("codigo")
	if cod == "" {
		http.Error(w, "Código não fornecido", http.StatusBadRequest)
		return
	}
	var item Iten
	// Busca o item onde o campo "codigo" é igual ao valor fornecido
	if err := bd.Where("codigo = ?", cod).First(&item).Error; err != nil {
		http.Error(w, "Item não encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

// Criar um novo item (envie JSON via POST)
func createItenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var item Iten
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Erro ao decodificar o item", http.StatusBadRequest)
		return
	}
	if err := bd.Create(&item).Error; err != nil {
		http.Error(w, "Erro ao criar o item", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(item)
}

// Atualizar um item (envie JSON via PUT, com o campo id preenchido)
func updateItenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var item Iten
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Erro ao decodificar o item", http.StatusBadRequest)
		return
	}
	if err := bd.Save(&item).Error; err != nil {
		http.Error(w, "Erro ao atualizar o item", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(item)
}

// Deletar um item (via query string: ?id=1)
func deleteItenHandler(w http.ResponseWriter, r *http.Request) {
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
	if err := bd.Delete(&Iten{}, id).Error; err != nil {
		http.Error(w, "Erro ao deletar o item", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Item deletado com sucesso"))
}

// ==================== HANDLERS PARA CATEGORIAS ====================

// Listar todas as categorias
func listCategoriasHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var cats []Cat
	if err := bd.Find(&cats).Error; err != nil {
		http.Error(w, "Erro ao buscar categorias", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cats)
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
	var cat Cat
	if err := bd.First(&cat, id).Error; err != nil {
		http.Error(w, "Categoria não encontrada", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(cat)
}

// Criar uma nova categoria (envie JSON via POST)
func createCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var cat Cat
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		http.Error(w, "Erro ao decodificar a categoria", http.StatusBadRequest)
		return
	}
	if err := bd.Create(&cat).Error; err != nil {
		http.Error(w, "Erro ao criar a categoria", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cat)
}

// Atualizar uma categoria (envie JSON via PUT, com o campo id preenchido)
func updateCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var cat Cat
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		http.Error(w, "Erro ao decodificar a categoria", http.StatusBadRequest)
		return
	}
	if err := bd.Save(&cat).Error; err != nil {
		http.Error(w, "Erro ao atualizar a categoria", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cat)
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
	if err := bd.Delete(&Cat{}, id).Error; err != nil {
		http.Error(w, "Erro ao deletar a categoria", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Categoria deletada com sucesso"))
}
