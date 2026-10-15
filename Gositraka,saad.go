package main

import (
	"fmt"
	"os"
	"strconv"
)

type ATHLETES struct { //Structure qui designe un athlete avec son nom, dossard et temps/
	nomcomplet string
	dossard    int
	temps      float64
}

const MAX_ATHLETES = 10 // La limite du tableau qui designe le nombre max d'athlete en constante/

var athlete [MAX_ATHLETES]ATHLETES //variable athlete qui est representer par un tableau avec la valeur du constante
var nombreAthletes int

func main() { //La page principale d'execution du programme ou on peut choisir/
	var choix int

	for {
		fmt.Println("====Menu Principal====")
		fmt.Println("1: Ajouter un Athlete")
		fmt.Println("2: Afficher la liste des resultats")
		fmt.Println("3: Rechercher un Athlete par son Numero de dossard")
		fmt.Println("4: Rechercher un Athletes par son Nom")
		fmt.Println("5: Afficher le classement")
		fmt.Println("6: Sauvegarder les résultats dans un fichier")
		fmt.Println("0: Quitter")
		fmt.Print("Votre choix : ")

		fmt.Scanln(&choix)

		if choix == 0 { //pour quitter notre programme*/
			fmt.Println("Au Revoir")
			break // pour stopper la boucle
		}

		switch choix {
		case 1:
			ajouter_athlete()
		case 2:
			afficherliste()
		case 3:
			rechercherpardossard()
		case 4:
			rechercherparnom()
		case 5:
			afficherClassement()
		case 6:
			sauvegarde()
		default:
			fmt.Println("Veuillez réessayer !!!")
		}
	}
}
func ajouter_athlete() { //fonction qui permet a l'utilisateur d'entrée le nom , performance et numero de dossard des athletes
	var doss int
	var nom string
	var perf float64

	fmt.Print("Nom complet:")
	fmt.Scanln(&nom)

	fmt.Print("Entrée le numero de Dossard:")
	fmt.Scanln(&doss)

	fmt.Print("Entrée sa performance en Seconde:")
	fmt.Scanln(&perf)

	athlete[nombreAthletes] = ATHLETES{dossard: doss, nomcomplet: nom, temps: perf}
	nombreAthletes = nombreAthletes + 1
	fmt.Println("---------------------------")
	fmt.Println("Athlète ajouté avec succès.")
	fmt.Println("---------------------------")

}

func afficherliste() { //Fonction pour afficher la liste des ATHLETES entrée par l'utilisateur /
	if nombreAthletes == 0 {
		fmt.Println("Aucun athlète enregistré.")
		return
	}

	fmt.Println("------ Liste des résultats ------")
	fmt.Println("Dossard", "Nom complet", "Performance (s)")
	fmt.Println("----------------------------------------------------------")
	for i := 0; i < nombreAthletes; i++ {
		fmt.Printf("%-5d %-10s %.2f\n", /* -5 pour decaler de 5 caractere a gauche , -10 pour decaller de la gauche a 10 caractere et .2 pour afficher 2 chiffres apres la virgules */
			athlete[i].dossard,
			athlete[i].nomcomplet,
			athlete[i].temps)
	}

}

func rechercherpardossard() { // fonction de recherche par dossard /
	var doss int
	fmt.Print("Entrez le numéro de dossard à rechercher : ")
	fmt.Scanln(&doss)
	for i := 0; i < nombreAthletes; i++ {
		if athlete[i].dossard == doss {
			fmt.Println("----------------------------------------------------------------------------")
			fmt.Println("Athlète trouvé :", athlete[i].nomcomplet, "Performance :", athlete[i].temps)
			fmt.Println("----------------------------------------------------------------------------")
		}

	}
}
func rechercherparnom() { // fonction de recherche par nom /
	var nom string
	fmt.Print("Entrez le nom complet à rechercher : ")
	fmt.Scanln(&nom)
	for i := 0; i < nombreAthletes; i++ {
		if athlete[i].nomcomplet == nom {
			fmt.Println("-----------------------------------------------------------------------------")
			fmt.Println("Athlète trouvé : Dossard", athlete[i].dossard, "Performance :", athlete[i].temps, "s")
			fmt.Println("-----------------------------------------------------------------------------")
		}

	}
}
func tri_par_rang() { //Tri a bulle qui consiste a comparer chaque rang de chaque athlete[i] par son successeur/
	for i := 0; i < nombreAthletes-1; i++ {
		for j := 0; j < nombreAthletes-i-1; j++ {
			if athlete[j].temps > athlete[j+1].temps {
				athlete[j], athlete[j+1] = athlete[j+1], athlete[j]
			}
		}
	}
}
func afficherClassement() { //affiche le classement de chaque athlete dans l'ordre et avec ex-aequo
	if nombreAthletes == 0 {
		fmt.Println("Aucun athlète enregistré.")
		return
	}
	//Pour gerer les ex-aequo , on crée 2 variables rang et compteur qui commence par 1 pour les comparer avec le reste/
	tri_par_rang()
	var rang int = 1
	var compteur int = 1

	fmt.Println("--- Classement des athlètes ---")
	fmt.Println("Rang", "", "Dossard", "", "Nom complet", "", "Performance (s)")
	fmt.Println("---------------------------------------------------------------------")

	for i := 0; i < nombreAthletes; i++ {
		if i > 0 && athlete[i].temps == athlete[i-1].temps {
			compteur = compteur + 1
		} else {
			rang = i + 1
			compteur = 1
		}

		fmt.Printf("%-5d %-10d %-25s %-15.2f\n",
			rang,
			athlete[i].dossard,
			athlete[i].nomcomplet,
			athlete[i].temps)
	}
}
func sauvegarde() { //permet de sauvegarder notre interaction avec la console en fichier txt
	file, err := os.OpenFile("resultats.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	// O_WRONLY pour ouvrir le fichier, O_CREATE pour crée le fichier s'il n'existe pas , O_TRUNC pour vider le fichier a l'ouverture
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer file.Close() // defer c'est une execution a la fin de la fonction

	file.WriteString("------ Liste des résultats ------\n")     //WriteString c'est pour ecrire des chaines dans un fichier
	file.WriteString("Dossard\tNom complet\tPerformance (s)\n") // \t =tabulation horizontal qui est equivaut a 8 chaines
	file.WriteString("----------------------------------------------------------\n")

	for i := 0; i < nombreAthletes; i++ {
		dossardconv := strconv.Itoa(athlete[i].dossard) // Pour convertir l'entier "dossard" en chaine de caractere*/
		tempsconv := strconv.FormatFloat(athlete[i].temps, 'f', 2, 64)

		ligne := dossardconv + "\t" + athlete[i].nomcomplet + "\t" + tempsconv + "\n"
		_, err := file.WriteString(ligne)
		if err != nil {
			fmt.Println("Erreur lors de l'écriture :", err)
			return
		}
	}

	fmt.Println("Les résultats ont été sauvegardés dans 'resultats.txt'")
}
