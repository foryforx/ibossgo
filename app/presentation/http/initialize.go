package http

import (
	"github.com/karuppaiah/ibossgo/app/application/presentation"
)

// UserHandler is the collection of all injection
type UserHandler struct {
	UserApplication presentation.IUserPresentation
}

// NewUserHandler will create new handler with all initialized
func NewUserHandler(
	userApp presentation.IUserPresentation,
) *UserHandler {
	return &UserHandler{
		UserApplication: userApp,
	}
}
