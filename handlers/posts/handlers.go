package posts

import (
	"net/http"
)

type Handler struct {

}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Post"))
}
