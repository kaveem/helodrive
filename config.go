package server

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	// SessionStore : store session globally
	SessionStore sessions.Store
	// OAuthConfig : store authConfigs globally
	OAuthConfig *oauth2.Config
	// Request : make http request global
	Request *http.Request
)

func init() {
	var cookieStore = sessions.NewCookieStore([]byte("something-very-secret"))
	cookieStore.Options = &sessions.Options{
		HttpOnly: true,
	}
	SessionStore = cookieStore

	// OAuthConfig = configureOAuthClient("404364039745-0caba0fvhaja2cogru4jvl0gqq3anf50.apps.googleusercontent.com", "zRly0iH-ThMZrYRxER5PT_ue")
	OAuthConfig = configureOAuthClient("404364039745-qe4mgo1jhqom3k7sve6usqh44c732u6c.apps.googleusercontent.com", "-WB8BpoMI4UZ3a675yJmXgjt")
	// OAuthConfig = configureOAuthClient("697255292974-h3ukerqst2o86h53dfbabikefm7dknm4.apps.googleusercontent.com", "DFhiPkfTF9yqwiqWGQlsrRFG")
}

func configureOAuthClient(clientID, clientSecret string) *oauth2.Config {
	redirectURL := os.Getenv("OAUTH2_CALLBACK")
	if redirectURL == "" {
		redirectURL = "http://localhost:3000/oauth2callback"
	}
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"email", "profile", "https://www.googleapis.com/auth/drive"},
		Endpoint:     google.Endpoint,
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
