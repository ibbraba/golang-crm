package main

import (
	"fmt"
)

func main() {

	clients := make(map[int]client)
	for {
		fmt.Println("Bienvenue sur Go CRM, que voulez-vous faire?")
		fmt.Println("1. Ajouter un client")
		fmt.Println("2. Lister les clients")
		fmt.Println("3. Modifier un client")
		fmt.Println("4. Supprimer un client")
		fmt.Println("5. Quitter")

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Nom du client:")
			var name string
			fmt.Scan(&name)
			fmt.Println("Email du client:")
			var email string
			fmt.Scan(&email)
			clients[len(clients)+1] = client{Nom: name, Email: email}
		case 2:
			fmt.Println("Liste des clients:")
			for id, c := range clients {
				fmt.Printf("ID: %d, Nom: %s, Email: %s\n", id, c.Nom, c.Email)
			}
		case 3:
			fmt.Println("ID du client à modifier:")
			var id int
			fmt.Scan(&id)
			fmt.Println("Nouveau nom du client:")
			var name string
			fmt.Scan(&name)
			fmt.Println("Nouvel email du client:")
			var email string
			fmt.Scan(&email)
		case 4:
			fmt.Println("ID du client à supprimer:")
			var id int
			fmt.Scan(&id)
		case 5:
			fmt.Println("Merci d'avoir utilisé Go CRM! A bientôt!")
			return
		default:
			fmt.Println("Commande invalide, veuillez réessayer.")
		}
	}
}

type client struct {
	Nom   string
	Email string
}
