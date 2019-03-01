package persistance

import (
	"github.com/jinzhu/gorm"
	"github.com/karuppaiah/ibossgo/app/domain/models"
	"github.com/pkg/errors"
	// log "github.com/sirupsen/logrus" // commented for easy toggle
	// "strconv" // commented for easy toggle
	// "strings" // commented for easy toggle
)

// User with Connection to DB
type User struct {
	Conn *gorm.DB
}

// SaveUser register a user so we know that we saw that user already.
func (p *User) SaveUser(u *models.User) error {
	isNewRecord := p.Conn.Debug().NewRecord(u)
	if isNewRecord {
		err := p.Conn.Debug().Create(&u).Error
		if err != nil {
			return errors.Wrapf(err, "postgres:SaveUser:error| user:%v", u)
		}
	}
	return nil
}

// LoadUser get data from a user.
func (p *User) LoadUserByEmail(email string) (result models.User, err error) {
	err = p.Conn.Debug().Where("email = ?", email).First(&result).Error
	if err != nil {
		return result, errors.Wrapf(err, "postgres:LoadUserByemail:error| email:%v", email)
	}
	return result, nil
}
