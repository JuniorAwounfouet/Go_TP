package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/JuniorAwounfouet/go-crm-app/internal/storage"
)

func Run(store storage.Storer) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Mini CRM v3!")

	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Add a contact")
		fmt.Println("2. List contacts")
		fmt.Println("3. Update a contact")
		fmt.Println("4. Delete a contact")
		fmt.Println("5. Exit")
		fmt.Print("Your choice: ")

		choice := readUserChoice(reader)

		switch choice {
		case 1:
			handleAddContact(reader, store)
		case 2:
			handleListContacts(store)
		case 3:
			handleUpdateContact(reader, store)
		case 4:
			handleDeleteContact(reader, store)
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, please try again")

		}
	}
}

// Les fonctions "handle..." s'occupent de l'interaction avec l'utilisateur
// et elles appellent la couche de stockage (store) pour effectuer les opérations.
// Elles sont découplées du stockage : elles fonctionnent avec n'importe quel storer

func handleAddContact(reader *bufio.Reader, storer storage.Storer) {
	fmt.Print("Enter contact name: ")
	name := readLine(reader)

	fmt.Print("Enter contact email: ")
	email := readLine(reader)

	contact := &storage.Contact{
		Name:  name,
		Email: email,
	}
	err := storer.Add(contact)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Contact '%s' added with ID %d.\n", contact.Name, contact.ID)
}

func handleListContacts(store storage.Storer) {
	contacts, err := store.GetAll()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(contacts) == 0 {
		fmt.Println(" No contacts to display.")
		return
	}

	fmt.Println("\n--- Contact List ---")
	for _, contact := range contacts {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", contact.ID, contact.Name, contact.Email)
	}
}

func handleUpdateContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Enter the ID of the contact to update: ")
	id := readInteger(reader)
	if id == -1 {
		return
	}

	// On vérifie que le contact existe avant de demander les nouvelles infos
	existingContact, err := store.GetByID(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Updating '%s'. Leave blank to keep current value.\n", existingContact.Name)

	fmt.Printf("New name (%s): ", existingContact.Name)
	newName := readLine(reader)

	fmt.Printf("New email (%s): ", existingContact.Email)
	newEmail := readLine(reader)

	err = store.Update(id, newName, newEmail)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Contact updated successfully.")
}

func handleDeleteContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Enter the ID of the contact to delete: ")
	id := readInteger(reader)
	if id == -1 {
		return
	}

	err := store.Delete(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Contact with ID %d has been deleted.\n", id)
} // Fonctions utilitaires pour la saisie utilisateur

func readLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readUserChoice(reader *bufio.Reader) int {
	choice, err := strconv.Atoi(readLine(reader))
	if err != nil {
		return -1 // Renvoie -1 pour un choix invalide
	}
	return choice
}

func readInteger(reader *bufio.Reader) int {
	id, err := strconv.Atoi(readLine(reader))
	if err != nil {
		fmt.Println("Error: Invalid ID. Please enter a number.")
		return -1
	}
	return id
}
