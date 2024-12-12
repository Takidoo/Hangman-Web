package main

import (
	"HangmanWeb/hangmanModule"
	"html/template"
	"net/http"
)

func acc(w http.ResponseWriter, r *http.Request) {
	tmps, _ := template.ParseFiles("templates/startup.html")
	tmps.Execute(w, hangmanModule.HangData)
}

func victoireHandle(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/victoire.html")
	tmpl.Execute(w, hangmanModule.HangData)
}

func defaiteHandle(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/defaite.html")
	hangmanModule.HangData = hangmanModule.PageData{
		WordToFind:     hangmanModule.Word,
		Try:            hangmanModule.Essais,
		Endd:           hangmanModule.Fin,
		Phrase:         hangmanModule.EndSentence,
		LettreEssayees: hangmanModule.Le,
	}
	tmpl.Execute(w, hangmanModule.HangData)
}

func scoreboardHandle(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/scoreboard.html")
	tmpl.Execute(w, hangmanModule.HangData)
}

func hangHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		hangmanModule.CheckLettre(r.FormValue("lettre"))
	}
	tmpl, _ := template.ParseFiles("templates/hangman.html")
	hangmanModule.HangData = hangmanModule.PageData{
		WordToFind:     string(hangmanModule.Rcw),
		Try:            hangmanModule.Essais,
		Endd:           hangmanModule.Fin,
		Phrase:         hangmanModule.EndSentence,
		LettreEssayees: hangmanModule.Le,
	}
	tmpl.Execute(w, hangmanModule.HangData)
}

func restart(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/reset.html")
	hangmanModule.InitGame()
	tmpl.Execute(w, hangmanModule.HangData)

}

func main() {
	hangmanModule.InitGame()
	http.HandleFunc("/scoreboard", scoreboardHandle)
	http.HandleFunc("/hangman", hangHandler)
	http.HandleFunc("/", acc)
	http.HandleFunc("/victoire", victoireHandle)
	http.HandleFunc("/defaite", defaiteHandle)
	http.HandleFunc("/reset", restart)
	http.Handle("/rscr/", http.StripPrefix("/rscr/", http.FileServer(http.Dir("./rscr"))))
	http.ListenAndServe(":80", nil)
}
