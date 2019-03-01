package connection

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // needed for gorm connection
	"github.com/karuppaiah/ibossgo/config"
	"github.com/pkg/errors"
)

// NewORMConn is to create a new GORM connection for workflow group package
func NewORMConn(conf *config.Configuration) (*gorm.DB, error) {
	var db *gorm.DB
	// If conf.Type is empty. Expected value postgresql
	if len(conf.Type) == 0 {
		return db, errors.New("conf.Type is empty")
	}
	// Validate other params
	if conf.Host == "" || conf.User == "" || conf.Name == "" {
		return db, errors.New("ConnectionString is invalid")
	}
	// SSL mode check for postgres
	if conf.SSLMode == "" {
		return db, errors.New("Postgress SSL mode not configured. Please add this to .env or env variable : PGSSLMODE=disable/enable")
	}
	if conf.Type == "postgresql" && (conf.SSLMode == "disable" || conf.SSLMode == "require") {
		db, err := gorm.Open("postgres", "host="+conf.Host+" port="+conf.Port+" user="+conf.User+" dbname="+conf.Name+" sslmode="+conf.SSLMode+" password="+conf.Password)
		if err != nil {
			return db, err
		}
		return db, err
	}
	return db, errors.Errorf("%v database is not supported(Type:Postgresql) or SSL mode not confiugured(Please add this to .env or env varibale : PGSSLMODE=disable/enable)", conf.Type)
}
