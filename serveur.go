package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type PageData struct {
	Title       string
	Player1     string
	Player2     string
	Grille      [][]string
	CurrentTurn string
	Message     string
}

func NouvelleGrille(lignes, colonnes int) [][]string {
	grille := make([][]string, lignes)
	for i := 0; i < lignes; i++ {
		ligne := make([]string, colonnes)
		for j := 0; j < colonnes; j++ {
			ligne[j] = ""
		}
		grille[i] = ligne
	}
	return grille
}

func Jouer(grille [][]string, col int, player string) {
	for i := len(grille) - 1; i >= 0; i-- {
		if grille[i][col] == "" {
			grille[i][col] = player
			break
		}
	}
}

func IsFull(grille [][]string) bool {
	for _, row := range grille {
		for _, cell := range row {
			if cell == "" {
				return false
			}
		}
	}
	return true
}

func VerifieVictoire(grille [][]string, player string) bool {
    lignes := len(grille)
    colonnes := len(grille[0])

    hasPion := func(i, j int) bool {
        if i < 0 || i >= lignes || j < 0 || j >= colonnes {
            return false
        }
        return grille[i][j] == player
    }

    for i := 0; i < lignes; i++ {
        for j := 0; j < colonnes; j++ {
            if hasPion(i, j) && hasPion(i, j+1) && hasPion(i, j+2) && hasPion(i, j+3) {
                return true
            }
            if hasPion(i, j) && hasPion(i+1, j) && hasPion(i+2, j) && hasPion(i+3, j) {
                return true
            }
            if hasPion(i, j) && hasPion(i+1, j+1) && hasPion(i+2, j+2) && hasPion(i+3, j+3) {
                return true
            }
            if hasPion(i, j) && hasPion(i-1, j+1) && hasPion(i-2, j+2) && hasPion(i-3, j+3) {
                return true
            }
        }
    }
    return false
}

func VerifieLigne(grille [][]string, symbole string) bool {
	lignes := len(grille)
	colonnes := len(grille[0])

	for i := 0; i < lignes; i++ {
		for j := 0; j < colonnes-3; j++ {
			if grille[i][j] == symbole &&
				grille[i][j+1] == symbole &&
				grille[i][j+2] == symbole &&
				grille[i][j+3] == symbole {
				return true
			}
		}
	}

	for i := 0; i < lignes-3; i++ {
		for j := 0; j < colonnes; j++ {
			if grille[i][j] == symbole &&
				grille[i+1][j] == symbole &&
				grille[i+2][j] == symbole &&
				grille[i+3][j] == symbole {
				return true
			}
		}
	}

	for i := 0; i < lignes-3; i++ {
		for j := 0; j < colonnes-3; j++ {
			if grille[i][j] == symbole &&
				grille[i+1][j+1] == symbole &&
				grille[i+2][j+2] == symbole &&
				grille[i+3][j+3] == symbole {
				return true
			}
		}
	}

	for i := 0; i < lignes-3; i++ {
		for j := 3; j < colonnes; j++ {
			if grille[i][j] == symbole &&
				grille[i+1][j-1] == symbole &&
				grille[i+2][j-2] == symbole &&
				grille[i+3][j-3] == symbole {
				return true
			}
		}
	}

	return false
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Page non trouvée", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, nil)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("contact.html")
	if err != nil {
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func Niveau(w http.ResponseWriter, r *http.Request, lignes, colonnes int, templateFile string) {
	player1 := r.URL.Query().Get("joueur1")
	player2 := r.URL.Query().Get("joueur2")
	currentTurn := r.URL.Query().Get("player")
	colStr := r.URL.Query().Get("col")

	grille := NouvelleGrille(lignes, colonnes)
	message := ""

	if colStr != "" && currentTurn != "" {
		col, err := strconv.Atoi(colStr)
		if err == nil && col >= 0 && col < colonnes {
			Jouer(grille, col, currentTurn)
      if VerifieVictoire(grille, currentTurn) {
    if currentTurn == "yellow" {
        message = "Le joueur JAUNE a gagné ! "
    } else {
        message = "Le joueur ROUGE a gagné !"
    }
}

			if currentTurn == "yellow" {
				currentTurn = "red"
			} else {
				currentTurn = "yellow"
			}
		}
	}

	data := PageData{
		Title:       "Puissance 4",
		Player1:     player1,
		Player2:     player2,
		Grille:      grille,
		CurrentTurn: currentTurn,
		Message:     message,
	}

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func Facile(w http.ResponseWriter, r *http.Request) {
	Niveau(w, r, 6, 7, "facile.html")
}

func Intermediaire(w http.ResponseWriter, r *http.Request) {
	Niveau(w, r, 6, 9, "intermediaire.html")
}

func Extreme(w http.ResponseWriter, r *http.Request) {
	Niveau(w, r, 7, 8, "extreme.html")
}

func Gagner(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("gagner.html")
	if err != nil {
		http.Error(w, "Page non trouvée", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, nil)
}

func Perdre(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("perdre.html")
	if err != nil {
		http.Error(w, "Page non trouvée", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/contact", Contact)
	http.HandleFunc("/facile", Facile)
	http.HandleFunc("/intermediaire", Intermediaire)
	http.HandleFunc("/extreme", Extreme)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Serveur lancé 🚀 sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
