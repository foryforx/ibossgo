package http

import (
	"github.com/jinzhu/gorm"
	"github.com/karuppaiah/ibossgo/app/user"
)

// UserHandler is the collection of all injection
type UserHandler struct {
	userDB user.User
}

// NewUserHandler will create new handler with all initialized
func NewUserHandler(
	dbConn *gorm.DB,
) *UserHandler {
	// DB initialization
	dbPort := user.NewUserRepo(dbConn)
	return &UserHandler{
		userDB: dbPort,
	}
}
