package application

import (
	"github.com/karuppaiah/ibossgo/app/application/persist"
	"github.com/karuppaiah/ibossgo/app/application/presentation"
	"github.com/karuppaiah/ibossgo/app/domain/models"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// UserApplication represents the application layer with persistance
type UserApplication struct {
	PersistDB persist.IUserPersistance
}

// NewUserApplication will create new user application with all initialized
func NewUserApplication(
	persistDB persist.IUserPersistance,
) presentation.IUserPresentation {
	return UserApplication{
		PersistDB: persistDB,
	}
}

// AuthUser authorize the user provided.
func (userApp UserApplication) AuthUser(u *models.User) (seen bool, err error) {
	if _, err = userApp.PersistDB.LoadUserByEmail(u.Email); err == nil {
		return true, nil
	} else {
		err = userApp.PersistDB.SaveUser(u)
		if err != nil {
			log.Errorln(err)
			return false, errors.Wrapf(err, "AuthUser: User cannot be saved")
		}
	}
	return false, nil
}
