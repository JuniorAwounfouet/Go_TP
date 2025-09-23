package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Structure Contact
type Contact struct {
	id    int
	nom   string
	email string
}

// Map globale des contacts, stocke des pointeurs vers Contact
var contacts = make(map[int]*Contact)

// Constructeur avec validation simple
func NewContact(id int, nom, email string) (*Contact, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID invalide")
	}
	if strings.TrimSpace(nom) == "" {
		return nil, fmt.Errorf("Nom vide")
	}
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") || strings.TrimSpace(email) == "" {
		return nil, fmt.Errorf("Email invalide")
	}
	return &Contact{id: id, nom: nom, email: email}, nil
}

// Affiche les informations du contact
func (c *Contact) afficher() {
	fmt.Printf("ID: %d, Nom: %s, Email: %s\n", c.id, c.nom, c.email)
}

// Met à jour le nom et/ou l'email du contact
func (c *Contact) mettreAJour(nom, email string) error {
	if nom != "" {
		c.nom = nom
	}
	if email != "" {
		if !strings.Contains(email, "@") || !strings.Contains(email, ".") || strings.TrimSpace(email) == "" {
			return fmt.Errorf("Email invalide")
		}
		c.email = email
	}
	return nil
}

// Lit une ligne depuis l'entrée utilisateur
func readLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Ajoute un contact via les flags CLI
func ajouterContactCLI(id int, nom, email string) {
	if _, exists := contacts[id]; exists {
		fmt.Println("Erreur : ID déjà utilisé.")
		return
	}
	contact, err := NewContact(id, nom, email)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	contacts[id] = contact
	fmt.Println("Contact ajouté")
}

// Ajoute un contact via le menu interactif
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

	contact, err := NewContact(id, nom, email)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	contacts[id] = contact
	fmt.Println("Contact ajouté")
}

// Liste tous les contacts
func listerContacts() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact trouvé")
		return
	}
	fmt.Println("Liste des contacts:")
	for _, c := range contacts {
		c.afficher()
	}
}

// Supprime un contact par ID
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

// Met à jour un contact existant
func mettreAJourContact(reader *bufio.Reader) {
	fmt.Print("Entrez ID du contact à mettre à jour: ")
	idStr := readLine(reader)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	contact, exists := contacts[id]
	if !exists {
		fmt.Println("Contact non trouvé.")
		return
	}

	fmt.Printf("Entrez nouveau Nom (actuel: %s, laisser vide pour ne pas changer): ", contact.nom)
	nom := readLine(reader)

	fmt.Printf("Entrez nouveau Email (actuel: %s, laisser vide pour ne pas changer): ", contact.email)
	email := readLine(reader)

	if err := contact.mettreAJour(nom, email); err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	fmt.Println("Contact mis à jour")
}

func main() {
	// Utilisation des flags pour ajout direct
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
