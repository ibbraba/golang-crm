package main

import (
	"flag"
	"fmt"
)

func main() {

	clients := make(map[int]client)

	nameFlag := flag.String("name", "", "Nom du client")
	mailFlag := flag.String("email", "", "Email du client")
	flag.Parse()

	// Gestion des flags
	if *nameFlag != "" && *mailFlag != "" {

		name := *nameFlag
		email := *mailFlag

		addClient(clients, name, email)
	}
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
			addClient(clients, name, email)
		case 2:
			listClients(clients)
		case 3:
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

			editClient(clients, id, name, email)

		case 4:
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

			deleteClient(clients, id)

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

func addClient(clients map[int]client, name string, email string) {
	clients[len(clients)+1] = client{Nom: name, Email: email}
	fmt.Println("Client ajouté avec succès! \n")
}

func listClients(clients map[int]client) {
	fmt.Println("Liste des clients:")

	// Si liste vide
	if len(clients) == 0 {
		fmt.Println("Aucun client enregistré. Veuillez en ajouter un. \n")

	}

	for id, client := range clients {
		fmt.Printf("ID: %d, Nom: %s, Email: %s\n", id, client.Nom, client.Email)
	}
}
func editClient(clients map[int]client, id int, name string, email string) {

	clients[id] = client{Nom: name, Email: email}
	fmt.Println("Client modifié avec succès! \n")
}

func deleteClient(clients map[int]client, id int) {

	delete(clients, id)
	fmt.Println("Client supprimé avec succès! \n")
}
