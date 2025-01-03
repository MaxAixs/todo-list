package handler

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	_ "todo-list/docs"
	"todo-list/todo/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) MapRoutes() http.Handler {
	router := mux.NewRouter()

	// Swagger router
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
		h.setupRoutes(lists, map[string]http.HandlerFunc{
			"create":  h.CreateList,
			"getAll":  h.GetAllLists,
			"getById": h.GetList,
			"update":  h.UpdateList,
			"delete":  h.DeleteList,
		})

		// Items routes within lists
		items := lists.PathPrefix("/{id}/items").Subrouter()
		h.setupRoutes(items, map[string]http.HandlerFunc{
			"create": h.CreateItem,
			"getAll": h.GetItems,
		})

		// Items routes
		itemsRoutes := api.PathPrefix("/items").Subrouter()
		h.setupRoutes(itemsRoutes, map[string]http.HandlerFunc{
			"getById":    h.GetItemById,
			"updateById": h.UpdateItem,
			"deleteById": h.DeleteItem,
		})
	}

	return router
}

func (h *Handler) setupRoutes(subRouter *mux.Router, handlers map[string]http.HandlerFunc) {
	subRouter.HandleFunc("/", handlers["create"]).Methods("POST")
	subRouter.HandleFunc("/", handlers["getAll"]).Methods("GET")
	subRouter.HandleFunc("/{id}", handlers["getById"]).Methods("GET")
	subRouter.HandleFunc("/{id}", handlers["updateById"]).Methods("PUT")
	subRouter.HandleFunc("/{id}", handlers["deleteById"]).Methods("DELETE")
}
