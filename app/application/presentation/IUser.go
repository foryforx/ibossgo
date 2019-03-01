package presentation

import (
	"github.com/karuppaiah/ibossgo/app/application/presentation/models"
)

// IUserPresentation is the list of ports allowed to be accessed by Domain layer for database
type IUserPresentation interface {
	SaveUser(u models.UserPresentation) error
	LoadUserByEmail(email string) (result models.UserPresentation, err error)
}
