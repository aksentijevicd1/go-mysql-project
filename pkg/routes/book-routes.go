package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/aksentijevicd1/go-mysql-project/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBooks = func(Router *mux.Router) {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	bh := controllers.NewBooks(l)

	Router.HandleFunc("/", bh.GetBooks).Methods(http.MethodGet)
	Router.HandleFunc("/{bookId}", bh.GetBookById).Methods(http.MethodGet)
	Router.HandleFunc("/{bookId}", bh.UpdateBook).Methods(http.MethodPut)
	Router.HandleFunc("/", bh.CreateBook).Methods(http.MethodPost)
	Router.HandleFunc("/{bookId}", bh.DeleteBook).Methods(http.MethodDelete)

}
