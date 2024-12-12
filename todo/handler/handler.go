package handler

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "todo-list/docs"
	"todo-list/todo/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) MapRoutes() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Authentication routes
	auth := router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/sign-up", h.signUp).Methods("POST")
	auth.HandleFunc("/sign-in", h.signIn).Methods("POST")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(h.AuthMiddleware)
	{
		// List routes
		lists := api.PathPrefix("/lists").Subrouter()
		lists.HandleFunc("/", h.CreateList).Methods("POST")
		lists.HandleFunc("/", h.GetAllLists).Methods("GET")
		lists.HandleFunc("/{id}", h.GetList).Methods("GET")
		lists.HandleFunc("/{id}", h.UpdateList).Methods("PUT")
		lists.HandleFunc("/{id}", h.DeleteList).Methods("DELETE")

		// Items routes within lists
		items := lists.PathPrefix("/{id}/items").Subrouter()
		items.HandleFunc("/", h.CreateItem).Methods("POST")
		items.HandleFunc("/", h.GetItems).Methods("GET")

		// Items routes
		itemsRoutes := api.PathPrefix("/items").Subrouter()
		itemsRoutes.HandleFunc("/{id}", h.GetItemById).Methods("GET")
		itemsRoutes.HandleFunc("/{id}", h.UpdateItem).Methods("PUT")
		itemsRoutes.HandleFunc("/{id}", h.DeleteItem).Methods("DELETE")
	}

	return router
}
