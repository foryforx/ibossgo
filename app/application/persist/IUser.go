package persist

import (
	"github.com/karuppaiah/ibossgo/app/domain/models"
)

// IUserPersistance is the list of ports allowed to be accessed by persistance layer
type IUserPersistance interface {
	SaveUser(u *models.User) error
	LoadUserByEmail(email string) (result models.User, err error)
}
