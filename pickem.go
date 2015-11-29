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

// var pickem = PickemLib.Data("PickemGameOne")
var whichGame = "first"
var whichWeek = "Week1"
var whichPlayer = "Jon"

var validPath = regexp.MustCompile("^/(pickem|other)/([a-zA-Z0-9]+)$")

func add(x, y int) int {
    return x + y
}

var templates = template.Must( template.New("").Funcs(template.FuncMap{ "add": add }).ParseGlob("templates/*") )

func renderTemplate(w http.ResponseWriter, tmpl string, p *PickemLib.Pickem) {

    err := templates.ExecuteTemplate(w, tmpl, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func pickemHandler(w http.ResponseWriter, r *http.Request) {

    var pickem = PickemLib.Saturate("PickemGame")
    renderTemplate(w, "pickem", pickem)
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

func picksHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "picks", nil)
}

func gameChoicesHandler(w http.ResponseWriter, r *http.Request, p *PickemLib.Pickem) {
    games := make([]PickemLib.Game, 0, len(p.PickemGames[whichGame].Weeks[whichWeek].Games) )
    for _, v := range p.PickemGames[whichGame].Weeks[whichWeek].Games { games = append( games, v ) }
    var choices = PickemLib.GamesChoices{ 
        Player : p.PickemGames[whichGame].Players[whichPlayer],
        Week : games,
    }
    if len(choices.Player.Picks) == 0 {
        for _, g := range choices.Week {
            choices.Player.Picks[g.Location] = PickemLib.GamesPicked{ Game : g, Points : 0 }
        }
    } 
    err := templates.ExecuteTemplate(w, "gameChoices", choices)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}

func interfaceHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "interface", nil)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, *PickemLib.Pickem)) http.HandlerFunc {

    var pickem = PickemLib.Saturate("PickemGame")
    return func(w http.ResponseWriter, r *http.Request) {
        /*  m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        } */
        fn(w, r, pickem)
    }
}

func main() {
	http.HandleFunc("/pickem/", pickemHandler)
	http.HandleFunc("/register/", registerHandler)
	http.HandleFunc("/leaderboard/", leaderboardHandler)
	http.HandleFunc("/privacy/", privacyHandler)
	http.HandleFunc("/donate/", donateHandler)
	http.HandleFunc("/contact/", contactHandler)
	http.HandleFunc("/picks/", picksHandler)
	http.HandleFunc("/gameChoices/", makeHandler(gameChoicesHandler))
	http.HandleFunc("/interface/", interfaceHandler)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
        	http.ServeFile(w, r, r.URL.Path[1:])
    		})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
