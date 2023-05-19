package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fbriandwi/databaseTest/models"
)

type ProductController struct {
	DB *sql.DB
}

func NewProductController(db *sql.DB) *ProductController {
	return &ProductController{
		DB: db,
	}
}

func (c *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	products, err := models.GetAllProducts(c.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to get products")
		log.Println(err)
		return
	}

	fmt.Println("All Products:")
	for _, p := range products {
		fmt.Fprintf(w, "ID: %d, Name: %s, Price: $%.2f\n", p.ID, p.Name, p.Price)
	}
}

func (c *ProductController) AddProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	idStr := r.FormValue("id")
	name := r.FormValue("name")
	priceStr := r.FormValue("price")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid price")
		return
	}

	product := models.Product{
		ID:    id,
		Name:  name,
		Price: price,
	}

	err = product.InsertProduct(c.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to add product")
		log.Println(err)
		return
	}

	fmt.Fprintf(w, "Product added successfully")
}

func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	err = models.DeleteProduct(c.DB, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to delete product")
		log.Println(err)
		return
	}

	fmt.Fprintf(w, "Product deleted successfully")
}
