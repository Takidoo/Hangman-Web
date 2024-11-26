package hangmanModule

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type PageData struct {
	WordToFind     string
	Try            int
	Endd           bool
	Phrase         string
	LettreEssayees string
}

var Word string
var CurrentWord string
var Rcw []rune
var HangData PageData
var Essais int
var Fin bool
var EndSentence string
var Le string

func InitWordList(file string) string {
	var l []string
	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		l = append(l, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return l[rand.Intn(len(l))]
}

func InitGame() {
	fmt.Printf("\x1bc")
	Le = ""
	Word = InitWordList(os.Args[1])
	Essais = 10
	Fin = false
	EndSentence, CurrentWord = "", ""
	for i := 0; i < len(Word); i++ {
		CurrentWord += "_"
	}
	Rcw = []rune(CurrentWord)
}

func CheckLettre(lettre string) {
	if Word == lettre {
		Fin = true
		EndSentence = "Bravo, vous avez gagné !"
	} else {
		Le += lettre + " "
		if strings.Contains(Word, lettre) {
			for i, l := range Word {
				if lettre == string(l) {
					Rcw[i] = l
				}
			}
		} else {
			Essais = Essais - 1
		}
	}
	if Essais == 0 {
		Fin = true
		EndSentence = "Dommage, vous avez perdu :("
	} else if string(Rcw) == Word {
		Fin = true
		EndSentence = "Bravo, vous avez gagné !"
	}
}
