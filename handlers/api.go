package handlers

import (
	"net/http"
)

type API interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	GetPost(w http.ResponseWriter, r *http.Request)
}
