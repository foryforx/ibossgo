package user

import (
	"github.com/jinzhu/gorm"
	"github.com/karuppaiah/ibossgo/app/models"
	"github.com/karuppaiah/ibossgo/app/user/repository"
)

// User is the list of ports allowed to be accessed by Domain layer for database
type User interface {
	SaveUser(u *models.User) error
	LoadUserByEmail(email string) (result models.User, err error)
}

// NewUserRepo To create new Repository with connection to Postgres
func NewUserRepo(conn *gorm.DB) User {
	return &repository.User{conn}
}
