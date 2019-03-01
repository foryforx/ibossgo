package presentation

import (
	"github.com/karuppaiah/ibossgo/app/domain/models"
)

// IUserPresentation is the list of ports allowed to be accessed by Domain layer for database
type IUserPresentation interface {
	AuthUser(u *models.User) (seen bool, err error)
}
