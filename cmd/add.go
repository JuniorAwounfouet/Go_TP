package cmd

import (
	"fmt"

	"minicrm/internal/storage"

	"github.com/spf13/cobra"
)

var addID int
var addNom string
var addEmail string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajouter un contact",
	RunE: func(cmd *cobra.Command, args []string) error {
		if addNom == "" || addEmail == "" {
			return fmt.Errorf("nom et email requis")
		}
		c := &storage.Contact{ID: addID, Name: addNom, Email: addEmail}
		if err := Store.Add(c); err != nil {
			return err
		}
		fmt.Printf("Contact ajouté (ID=%d)\n", c.ID)
		return nil
	},
}

func init() {
	addCmd.Flags().IntVar(&addID, "id", 0, "ID (optionnel)")
	addCmd.Flags().StringVar(&addNom, "nom", "", "Nom du contact")
	addCmd.Flags().StringVar(&addEmail, "email", "", "Email du contact")
	RootCmd.AddCommand(addCmd)
}
