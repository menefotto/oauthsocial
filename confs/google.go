package conf

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// NewGoogle creates a new oauth2 configuration suitable to talk to
// google oauth provider,email as a scope must be there
func NewGoogle(id, secret, callback string, scopes []string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		RedirectURL:  callback,
		Scopes:       scopes,
		Endpoint:     google.Endpoint,
	}
}

const GoogleUrl = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
