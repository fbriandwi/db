	package routers

	import (
		"github.com/gorilla/mux"
		"github.com/fbriandwi/databaseTest/controllers"
	)

	func NewRouter(productController *controllers.ProductController) *mux.Router {
		router := mux.NewRouter()

		router.HandleFunc("/products", productController.GetAllProducts).Methods("GET")
		router.HandleFunc("/products", productController.AddProduct).Methods("POST")
		router.HandleFunc("/products", productController.DeleteProduct).Methods("DELETE")

		return router
	}
