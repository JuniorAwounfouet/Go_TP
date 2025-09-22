package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var contacts = make(map[int][]string)

func readLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Fonction pour ajouter un contact via les flags CLI
func ajouterContactCLI(id int, nom, email string) {
	if _, exists := contacts[id]; exists {
		fmt.Println("Erreur : ID déjà utilisé.")
		return
	}
	contacts[id] = []string{nom, email}
	fmt.Println("Contact ajouté")
}

// Fonction pour ajouter un contact
func ajouterContact(reader *bufio.Reader) {
	fmt.Print("Entrez ID: ")
	idStr := readLine(reader)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID invalide.")
		return
	}
	if _, exists := contacts[id]; exists {
		fmt.Println("Erreur : ID déjà utilisé")
		return
	}
	fmt.Print("Entrez le nom: ")
	nom := readLine(reader)
	fmt.Print("Entrez l'email: ")
	email := readLine(reader)

	contacts[id] = []string{nom, email}
	fmt.Println("Contact ajouté")
}

// Fonction pour lister les contacts
func listerContacts() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact trouvé")
		return
	}
	fmt.Println("Liste des contacts:")
	for id, info := range contacts {
		fmt.Printf("ID: %d, Nom: %s, Email: %s\n", id, info[0], info[1])
	}
}

// Fonction pour supprimer un contact
func supprimerContact(reader *bufio.Reader) {
	fmt.Print("Entrez ID à supprimer: ")
	idStr := readLine(reader)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	if _, exists := contacts[id]; exists {
		delete(contacts, id)
		fmt.Println("Contact supprimé")
	} else {
		fmt.Println("Contact non trouvé")
	}
}

// Fonction pour mettre à jour un contact
func mettreAJourContact(reader *bufio.Reader) {
	fmt.Print("Entrez ID du contact à mettre à jour: ")
	idStr := readLine(reader)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	info, exists := contacts[id]
	if !exists {
		fmt.Println("Contact non trouvé.")
		return
	}

	fmt.Printf("Entrez nouveau Nom (actuel: %s): ", info[0])
	nom := readLine(reader)

	fmt.Printf("Entrez nouveau Email (actuel: %s): ", info[1])
	email := readLine(reader)

	contacts[id] = []string{nom, email}
	fmt.Println("Contact mis à jour")
}

func main() {

	flagAjouter := flag.Bool("ajouter", false, "Ajouter un contact")
	flagID := flag.Int("id", 0, "ID du contact")
	flagNom := flag.String("nom", "", "Nom du contact")
	flagEmail := flag.String("email", "", "Email du contact")
	flag.Parse()

	if *flagAjouter {
		if *flagID == 0 || *flagNom == "" || *flagEmail == "" {
			fmt.Println("Pour ajouter un contact, veuillez fournir id, nom et email")
			return
		}
		ajouterContactCLI(*flagID, *flagNom, *flagEmail)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n** Gestionnaire de Contacts **")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Mettre à jour un contact")
		fmt.Println("5. Quitter")

		fmt.Print("Choisissez une option: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			ajouterContact(reader)
		case "2":
			listerContacts()
		case "3":
			supprimerContact(reader)
		case "4":
			mettreAJourContact(reader)
		case "5":
			fmt.Println("Au revoir :)")
			return
		default:
			fmt.Println("Option invalide, veuillez réessayer.")
		}
	}
}
