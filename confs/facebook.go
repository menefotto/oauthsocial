package conf

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

// NewFacebook creates a new oauth2 configuration suitable to talk to
// facebook oauth provider,email as a scope must be there
func NewFacebook(id, secret, callback string, scopes []string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		RedirectURL:  callback,
		Scopes:       scopes,
		Endpoint:     facebook.Endpoint,
	}
}

// Facebook Url requires to add the fields of interest before using it otherwi// otherwise you just get back Id and first name last name not even the email/// you should add the following "&fields=id,name,picture" etc consult the face// book developers guide at https://developers.facebook.com/docs/facebook-login/permissions
const FacebookUrl = "https://graph.facebook.com/me?access_token="
