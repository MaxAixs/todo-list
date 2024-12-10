package handler

import (
	"net/http"
)

func (h *Handler) CreateList(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) GetList(w http.ResponseWriter, r *http.Request)     {}
func (h *Handler) UpdateList(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) DeleteList(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) GetAllLists(w http.ResponseWriter, r *http.Request) {}
