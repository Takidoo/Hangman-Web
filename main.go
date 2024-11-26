package main

import (
	"HangmanWeb/hangmanModule"
	"html/template"
	"net/http"
)

func hangHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		hangmanModule.CheckLettre(r.FormValue("lettre"))
	}
	tmpl, _ := template.ParseFiles("hangman.html")
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
	tmpl, _ := template.ParseFiles("reset.html")
	hangmanModule.InitGame()
	tmpl.Execute(w, hangmanModule.HangData)

}

func main() {
	hangmanModule.InitGame()
	http.HandleFunc("/", hangHandler)
	http.HandleFunc("/reset", restart)
	http.ListenAndServe(":80", nil)
}
