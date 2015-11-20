// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"
    "regexp"
    "html/template"
    "github.com/Comp-698/PickemLib"
)

type Pickem struct {
    Picks map[string]int
    Name string
}

var validPath = regexp.MustCompile("^/(pickem|other)/([a-zA-Z0-9]+)$")

// var templates = template.Must( template.ParseFiles( "pickem.html" ) )
var templates = template.Must(template.ParseGlob("templates/*"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Pickem) {
    err := templates.ExecuteTemplate(w, tmpl, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func pickemHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "pickem", &Pickem{Name: "pickem"})
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "register", nil)
}

func privacyHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "privacy", nil)
}

func leaderboardHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "leaderboard", nil)
}

func donateHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "donate", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "contact", nil)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}

func main() {
	http.HandleFunc("/pickem/", pickemHandler)
	http.HandleFunc("/register/", registerHandler)
	http.HandleFunc("/leaderboard/", leaderboardHandler)
	http.HandleFunc("/privacy/", privacyHandler)
	http.HandleFunc("/donate/", donateHandler)
	http.HandleFunc("/contact/", contactHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte( PickemLib.Data("paste") ) )
}
