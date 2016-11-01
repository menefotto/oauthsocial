package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/oauthsocial"
)

func main() {

	social := oauthsocial.New(confs.NewGoogle("YourId", "YourSecret", "callbackAddress", []string{"mail"}), confs.GoogleUrl, OnDataExample)

	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)
	mux.HandleFunc("/auth/google", social.Handler)
	mux.HandleFunc("/auth/google/callback", social.Callback)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h2> Oauth social example! </h2>")
	return
}

func OnDataExample(response string) (interface{}, error) {
	log.Println("got response")
	// do something with it
	return
}
