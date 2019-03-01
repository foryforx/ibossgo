package models

import (
	"github.com/jinzhu/gorm"
	// "github.com/pkg/errors" // commented for easy toggle
	// log "github.com/sirupsen/logrus" // commented for easy toggle
	// "strconv" // commented for easy toggle
	// "strings" // commented for easy toggle
)

// User is a retrieved and authentiacted user.
type User struct {
	gorm.Model
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}
