package handler

import (
	"net/http"
)

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) GetItem(w http.ResponseWriter, r *http.Request)    {}
func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {}
