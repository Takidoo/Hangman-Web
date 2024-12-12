package Json

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Score struct {
	Username string `json:"username"`
	Score    int    `json:"score"`
}

func ajout(file string) {
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
	return
}

func readJSON(filePath string) ([]Score, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var scores []Score
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&scores)
	if err != nil {
		return nil, err
	}

	return scores, nil
}

func writeJSON(filePath string, scores []Score) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(scores)
	if err != nil {
		return err
	}

	return nil
}

func addScore(scores []Score, username string, score int) []Score {
	newScore := Score{
		Username: username,
		Score:    score,
	}
	scores = append(scores, newScore)
	return scores
}

func main() {
	filePath := "rscr/scores.json"

	scores, err := readJSON(filePath)
	if err != nil {
		log.Fatal("Erreur lors de la lecture du fichier JSON:", err)
	}

	scores = addScore(scores, "Joueur4", 0)

	err = writeJSON(filePath, scores)
	if err != nil {
		log.Fatal("Erreur lors de l'écriture dans le fichier JSON:", err)
	}

	fmt.Println("\nLe score a été ajouté et les données sont sauvegardées dans le fichier.")
}
