package main

import (
	"flag"
	"fmt"
)

func main() {
	action := flag.String("action", "", "rechercher, ajouter, supprimer, modifier, lister")
	nom := flag.String("nom", "", "Nom de l'utilisateur")
	prenom := flag.String("prenom", "", "Prénom de l'utilisateur")
	numero := flag.String("numero", "", "Numéro de téléphone")
	prenomModif := flag.String("prenomModif", "", "Prénom de l'utilisateur à changer")

	flag.Parse()
	switch *action {
	case "ajouter":
		addUser(*prenom, *nom, *numero)
	case "supprimer":
		suppUser(*prenom)
	case "modifier":
		modify(*prenom, *nom, *numero, *prenomModif)
	case "lister":
		listAnnuaire()
	case "rechercher":
		search(*prenom)
	default:
		fmt.Println("option non reconnu")
	}
}

type annuaire struct {
	nom, prenom string
	phoneNumber string
}

var annuaireArray = map[string]annuaire{
	"Elwin": annuaire{
		nom:         "Bassaget",
		prenom:      "Elwin",
		phoneNumber: "0630416103",
	},
	"Joseph": annuaire{
		nom:         "Joestar",
		prenom:      "Joseph",
		phoneNumber: "0632450231",
	},
}

func search(nom string) annuaire {

	fmt.Println(annuaireArray[nom].nom)
	fmt.Println(annuaireArray[nom].prenom)
	fmt.Println(annuaireArray[nom].phoneNumber)

	entry := annuaireArray[nom]
	return entry

}

func addUser(prenom string, nom string, numero string) annuaire {
	_, exist := annuaireArray[prenom]
	if !exist {
		annuaireArray[prenom] = annuaire{
			nom:         nom,
			prenom:      prenom,
			phoneNumber: numero,
		}
		fmt.Println("Nouvelle addition ajouté à l'annuaire ")
		listAnnuaire()
		return annuaireArray[prenom]
	}

	fmt.Println("Personne existe déjà ")
	return annuaireArray[prenom]

}

func listAnnuaire() {
	var keys []string
	for k := range annuaireArray {
		keys = append(keys, k)
	}
	for _, k := range keys {
		v := annuaireArray[k]
		fmt.Printf("%s\n%s\n%s\n \n", v.prenom, v.nom, v.phoneNumber)
	}
}

func suppUser(prenom string) {
	delete(annuaireArray, prenom)
	listAnnuaire()
}

func modify(prenom string, nom string, numero string, prenomModif string) annuaire {
	var newPrenom string

	if prenomModif == "" {
		newPrenom = prenom
	} else {
		newPrenom = prenomModif
	}

	modif := annuaireArray[prenom]
	if nom != "" {
		modif.nom = nom

	}
	if numero != "" {
		modif.phoneNumber = numero

	}
	modif.prenom = newPrenom
	annuaireArray[newPrenom] = modif

	return annuaireArray[newPrenom]
}
