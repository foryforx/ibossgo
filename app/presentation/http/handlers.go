package http

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/karuppaiah/ibossgo/app/domain/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

var cred Credentials
var conf *oauth2.Config

// Credentials which stores google ids.
type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

// RandToken generates a random @l length token.
func RandToken(l int) string {
	b := make([]byte, l)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}

func getLogoutURL() string {
	return "http://127.0.0.1:9090/logout"
}

func init() {
	file, err := ioutil.ReadFile("./creds.json")
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &cred)

	conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  "http://127.0.0.1:9090/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}
}

// IndexHandler handels /.
func (h *UserHandler) IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

// AuthHandler handles authentication of a user and initiates a session.
func (h *UserHandler) AuthHandler(c *gin.Context) {
	// Handle the exchange code to initiate a transport.
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	queryState := c.Request.URL.Query().Get("state")
	if retrievedState != queryState {
		link := getLogoutURL()
		log.Printf("Invalid session state: retrieved: %s; Param: %s", retrievedState, queryState)
		c.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{"message": "Invalid session state.", "link": link})
		return
	}
	code := c.Request.URL.Query().Get("code")
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Login failed. Please try again."})
		return
	}
	session.Set("AccessToken", tok.AccessToken)
	session.Set("RefreshToken", tok.RefreshToken)
	session.Set("TokenType", tok.TokenType)
	session.Set("Expiry", tok.Expiry.Format(time.RFC3339))
	session.Save()
	client := conf.Client(oauth2.NoContext, tok)
	userinfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	defer userinfo.Body.Close()
	data, _ := ioutil.ReadAll(userinfo.Body)
	u := models.User{}
	if err = json.Unmarshal(data, &u); err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error marshalling response. Please try agian."})
		return
	}
	session.Set("user-id", u.Email)
	err = session.Save()
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error while saving session. Please try again."})
		return
	}
	seen := false
	seen, err = h.UserApplication.AuthUser(&u)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error while saving user. Please try again."})
		return
	}
	link := getLogoutURL()
	c.HTML(http.StatusOK, "battle.tmpl", gin.H{"email": u.Email, "seen": seen, "link": link})
}

// LoginHandler handles the login procedure.
func (h *UserHandler) LoginHandler(c *gin.Context) {
	state := RandToken(32)
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()
	link := getLoginURL(state)
	c.HTML(http.StatusOK, "auth.tmpl", gin.H{"link": link})
}

// FieldHandler is a rudementary handler for logged in users.
func (h *UserHandler) FieldHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "field.tmpl", gin.H{"user": userID})
}

// LogoutHandler handles logout calls
func (h *UserHandler) LogoutHandler(c *gin.Context) {
	logout(c)
	c.HTML(http.StatusOK, "logout.tmpl", gin.H{"link": "http://127.0.0.1:9090/login"})
}

func logout(c *gin.Context) {
	s := sessions.Default(c)
	s.Delete("user-id")
	s.Save()
}
func extractPath(next string) string {
	n, err := url.Parse(next)
	if err != nil {
		return "/"
	}
	return n.Path
}
