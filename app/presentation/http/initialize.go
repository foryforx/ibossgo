package http

import (
	"github.com/jinzhu/gorm"
	"github.com/karuppaiah/ibossgo/app/application/presentation"
)

// NewUserPresentation To create new Repository with connection to Postgres
func NewUserPresentation(conn *gorm.DB) presentation.IUserPresentation {
	return &User{conn}
}
