package users

import (
	"net/http"
)

type Handler struct {

}

func (u *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}
