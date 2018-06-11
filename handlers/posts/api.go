package posts

import (
	"net/http"
)

type PostRoutes interface {
	GetPost(w http.ResponseWriter, r *http.Request)
}
