package todo

import (
	"github.com/gorilla/mux"
)

type Handler struct{}

func (h *Handler) MapRoutes() *mux.Router {
	router := mux.NewRouter()

	// Authentication routes
	router.HandleFunc("/auth/sign-up", h.signUp).Methods("POST")
	router.HandleFunc("/auth/sign-in", h.signIn).Methods("POST")

	// List routes
	router.HandleFunc("/lists", h.CreateList).Methods("POST")
	router.HandleFunc("/lists", h.GetAllLists).Methods("GET")
	router.HandleFunc("/lists/{id}", h.GetList).Methods("GET")
	router.HandleFunc("/lists/{id}", h.UpdateList).Methods("PUT")
	router.HandleFunc("/lists/{id}", h.DeleteList).Methods("DELETE")

	// Item routes
	router.HandleFunc("/lists/{id}/items", h.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", h.GetItem).Methods("GET")
	router.HandleFunc("/items/{id}", h.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", h.DeleteItem).Methods("DELETE")

	return router
}
