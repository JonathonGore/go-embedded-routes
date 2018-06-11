package users

import (
	"net/http"
)

type UserRoutes interface {
	GetUser(w http.ResponseWriter, r *http.Request)
}
