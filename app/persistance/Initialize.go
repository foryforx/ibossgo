package persistance

import (
	"github.com/jinzhu/gorm"
	"github.com/karuppaiah/ibossgo/app/application/persist"
)

// NewUserStore To create new Repository with connection to Postgres
func NewUserStore(conn *gorm.DB) persist.IUserPersistance {
	return &User{conn}
}
