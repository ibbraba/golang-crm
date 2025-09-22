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
			fmt.Println("Client ajouté avec succès! \n")
		case 2:
			fmt.Println("Liste des clients:")

			// Si liste vide
			if len(clients) == 0 {
				fmt.Println("Aucun client enregistré. Veuillez en ajouter un. \n")
				continue
			}

			for id, client := range clients {
				fmt.Printf("ID: %d, Nom: %s, Email: %s\n", id, client.Nom, client.Email)
			}
		case 3:

			// Si liste vide
			if len(clients) == 0 {
				fmt.Println("Aucun client enregistré. Veuillez en ajouter un. \n")
				continue
			}

			fmt.Println("ID du client à modifier:")
			var id int
			fmt.Scan(&id)

			// Vérifie si le client existe
			if _, exists := clients[id]; !exists {
				fmt.Printf("Client avec l'ID %d non trouvé. \n", id)
				continue
			}

			fmt.Println("Nouveau nom du client:")
			var name string
			fmt.Scan(&name)
			fmt.Println("Nouvel email du client:")
			var email string
			fmt.Scan(&email)

			clients[id] = client{Nom: name, Email: email}
			fmt.Println("Client modifié avec succès! \n")

		case 4:
			// Si liste vide
			if len(clients) == 0 {
				fmt.Println("Aucun client enregistré. Veuillez en ajouter un. \n")
				continue
			}

			fmt.Println("ID du client à supprimer:")
			var id int
			fmt.Scan(&id)
			if _, exists := clients[id]; !exists {
				fmt.Printf("Client avec l'ID %d non trouvé. \n", id)
				continue
			}

			delete(clients, id)
			fmt.Println("Client supprimé avec succès! \n")

		case 5:
			fmt.Println("Merci d'avoir utilisé Go CRM! A bientôt!")
			return
		default:
			fmt.Println("Commande invalide, veuillez réessayer. \n")
		}

		fmt.Println()
	}
}

type client struct {
	Nom   string
	Email string
}
