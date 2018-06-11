package handlers

import (
	"github.com/JonathonGore/go-embedded-routes/handlers/users"
	"github.com/JonathonGore/go-embedded-routes/handlers/posts"
)

type Handler struct {
	users.UserRoutes
	posts.PostRoutes
}

func New() (Handler, error) {
	handler := Handler{&users.Handler{}, &posts.Handler{}}

	return handler, nil
}
