package main

import (
	"fmt"
)

const (
	Lignes  = 6
	Colonnes = 7
)

func main() {
	// Créer la grille
	grille := make([][]string, Lignes)
	for i := 0; i < Lignes; i++ {
		grille[i] = make([]string, Colonnes)
		for j := 0; j < Colonnes; j++ {
			grille[i][j] = "." // "." = case vide
		}
	}

	// Afficher la grille
	afficherGrille(grille)
}

func afficherGrille(grille [][]string) {
	for i := 0; i < len(grille); i++ {
		for j := 0; j < len(grille[i]); j++ {
			fmt.Print(grille[i][j], " ")
		}
		fmt.Println()
	}
}
