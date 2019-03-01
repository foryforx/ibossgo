package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/karuppaiah/ibossgo/app/application"
	"github.com/karuppaiah/ibossgo/app/persistance"
	"github.com/karuppaiah/ibossgo/app/presentation/http"
	"github.com/karuppaiah/ibossgo/config"
	"github.com/karuppaiah/ibossgo/connection"
	"github.com/karuppaiah/ibossgo/middleware"
	log "github.com/sirupsen/logrus"
	"os"
)

var isProduction bool

func init() {
	// Init env variables
	if os.Getenv("ENV") == "production" {
		isProduction = true
	}
}
func init() {

}

func main() {
	// Creating global configuration object
	configuration := config.GetConfiguration(isProduction)
	configuration.Print()

	// Dependency injection
	// GORM connection
	db, dbErr := connection.NewORMConn(configuration)
	if dbErr != nil {
		log.Fatalf("Error creating connType with ORM: %v", dbErr)
	}
	defer db.Close()
	userPersistance := persistance.NewUserStore(db)
	userApplication := application.NewUserApplication(userPersistance)
	userHandler := http.NewUserHandler(userApplication)

	// router setup
	router := gin.Default()
	store := sessions.NewCookieStore([]byte(http.RandToken(64)))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(sessions.Sessions("goquestsession", store))
	router.Static("/css", "./ui/static/css")
	router.Static("/img", "./ui/static/img")
	router.LoadHTMLGlob("ui/templates/*")

	router.GET("/", userHandler.IndexHandler)
	router.GET("/login", userHandler.LoginHandler)
	router.GET("/auth", userHandler.AuthHandler)
	router.GET("/logout", userHandler.LogoutHandler)
	authorized := router.Group("/battle")
	authorized.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/field", userHandler.FieldHandler)
	}

	router.Run("127.0.0.1:9090")
}
