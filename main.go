package main

import (
	"flag"
	"fmt"
)

func main() {

	contacts := make(map[int]*contact)

	nameFlag := flag.String("name", "", "Nom du contact")
	mailFlag := flag.String("email", "", "Email du contact")
	flag.Parse()

	// Gestion des flags
	if *nameFlag != "" && *mailFlag != "" {

		name := *nameFlag
		email := *mailFlag

		c := NewContact(name, email)
		c.addcontact(contacts, c.Nom, c.Email)
	}
	for {
		fmt.Println("Bienvenue sur Go CRM, que voulez-vous faire?")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister les contacts")
		fmt.Println("3. Modifier un contact")
		fmt.Println("4. Supprimer un contact")
		fmt.Println("5. Quitter")

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Nom du contact:")
			var name string
			fmt.Scan(&name)
			fmt.Println("Email du contact:")
			var email string
			fmt.Scan(&email)
			c := NewContact(name, email)
			c.addcontact(contacts, c.Nom, c.Email)

		case 2:
			listcontacts(contacts)
		case 3:
			if len(contacts) == 0 {
				fmt.Println("Aucun contact enregistré. Veuillez en ajouter un. \n")
				continue
			}

			fmt.Println("ID du contact à modifier:")
			var id int
			fmt.Scan(&id)

			// Vérifie si le contact existe
			if _, exists := contacts[id]; !exists {
				fmt.Printf("contact avec l'ID %d non trouvé. \n", id)
				continue
			}

			fmt.Println("Nouveau nom du contact:")
			var name string
			fmt.Scan(&name)
			fmt.Println("Nouvel email du contact:")
			var email string
			fmt.Scan(&email)
			c := contact{Nom: name, Email: email}
			c.editcontact(contacts, id, name, email)

		case 4:
			if len(contacts) == 0 {
				fmt.Println("Aucun contact enregistré. Veuillez en ajouter un. \n")
				continue
			}
			fmt.Println("ID du contact à supprimer:")
			var id int
			fmt.Scan(&id)

			if _, exists := contacts[id]; !exists {
				fmt.Printf("contact avec l'ID %d non trouvé. \n", id)
				continue
			}
			c := contact{}
			c.deletecontact(contacts, id)

		case 5:
			fmt.Println("Merci d'avoir utilisé Go CRM! A bientôt!")
			return
		default:
			fmt.Println("Commande invalide, veuillez réessayer. \n")
		}

		fmt.Println()
	}
}

type contact struct {
	Nom   string
	Email string
}

func NewContact(nom string, email string) *contact {
	return &contact{
		Nom:   nom,
		Email: email,
	}
}

func (c contact) addcontact(contacts map[int]*contact, name string, email string) {
	contacts[len(contacts)+1] = &contact{Nom: name, Email: email}
	fmt.Println("contact ajouté avec succès! \n")
}

func listcontacts(contacts map[int]*contact) {
	fmt.Println("Liste des contacts:")

	// Si liste vide
	if len(contacts) == 0 {
		fmt.Println("Aucun contact enregistré. Veuillez en ajouter un. \n")

	}

	for id, contact := range contacts {
		fmt.Printf("ID: %d, Nom: %s, Email: %s\n", id, contact.Nom, contact.Email)
	}
}
func (c contact) editcontact(contacts map[int]*contact, id int, name string, email string) {

	contacts[id] = &contact{Nom: name, Email: email}
	fmt.Println("contact modifié avec succès! \n")
}

func (c contact) deletecontact(contacts map[int]*contact, id int) {

	delete(contacts, id)
	fmt.Println("contact supprimé avec succès! \n")
}
