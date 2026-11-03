package routes

import (
	// "github.com/samson4/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
	"github.com/samson4/go-bookstore/pkg/controllers"
)

func RegisterBookStoreRoutes(router *mux.Router){
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBookById).Methods("DELETE")
}