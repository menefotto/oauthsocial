package oauthsocial

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

func New(conf *oauth2.Config, url string, f OnData) *SocialOauth {
	return &SocialOauth{conf, url, f}
}

type SocialOauth struct {
	conf    *oauth2.Config
	url     string
	functor OnData
}

type OnData func(reponse []byte) (interface{}, error)

func (o *SocialOauth) Handler(w http.ResponseWriter, r *http.Request) {
	url := gConf.AuthCodeURL("state", oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (o *SocialOauth) Callback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := o.conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	resp, err := http.Get(o.url + url.QueryEscape(token.AccessToken))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	fields := &FbResponse{}
	err = json.Unmarshal(response, fields)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
	}

	o.functor(response)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

}
