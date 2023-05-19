package models

import (
	"database/sql"
	"log"
)

type Product struct {
	ID    int
	Name  string
	Price float64
}

func GetAllProducts(db *sql.DB) ([]Product, error) {
	query := "SELECT id, name, price FROM products"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			log.Println(err)
			continue
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (p *Product) InsertProduct(db *sql.DB) error {
	query := "INSERT INTO products (id, name, price) VALUES (?, ?, ?)"
	_, err := db.Exec(query, p.ID, p.Name, p.Price)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(db *sql.DB, id int) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
