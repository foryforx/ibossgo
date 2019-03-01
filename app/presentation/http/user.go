package http

import (
	"github.com/jinzhu/gorm"
	"github.com/karuppaiah/ibossgo/app/application/presentation"
	"github.com/karuppaiah/ibossgo/app/persistance"
)

// UserHandler is the collection of all injection
type UserHandler struct {
	userApplication presentation.IUserPresentation
}

// NewUserHandler will create new handler with all initialized
func NewUserHandler(
	userApp presentation.IUserPresentation
) *UserHandler {
	return &UserHandler{
		userApplication: userApp,
	}
}
