package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/fbriandwi/databaseTest/controllers"
	"github.com/fbriandwi/databaseTest/routers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "products.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	controllers.CreateTables(db)

	productController := controllers.NewProductController(db)

	router := routers.NewRouter(productController)
	log.Fatal(http.ListenAndServe(":8080", router))
}
