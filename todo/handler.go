package todo

import "net/http"

type Handler struct {
}

func (h *Handler) MapsRoutes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/auth/sign-up", h.signUp) //POST
	router.HandleFunc("/auth/sign-in", h.signIn) // POST

	router.HandleFunc("/lists", h.CreateList)     //POST
	router.HandleFunc("/lists/", h.GetAllLists)   // GET getAllLists
	router.HandleFunc("/lists/:id", h.GetList)    // GET getById
	router.HandleFunc("/lists/:id", h.UpdateList) // PUT updateById
	router.HandleFunc("/lists/:id", h.DeleteList) // DEL deleteByID

	router.HandleFunc("/lists/:id/items", h.CreateItem) // POST createItem

	router.HandleFunc("/items/:id", h.GetItem)    //GET getItemsById
	router.HandleFunc("/items/:id", h.UpdateItem) // PUT UpdateItemsById
	router.HandleFunc("/items/:id", h.DeleteItem) // DEL deleteItemsById

	return router
}
