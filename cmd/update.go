package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updID int
var updName string
var updEmail string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Mettre à jour un contact",
	RunE: func(cmd *cobra.Command, args []string) error {
		if updID == 0 {
			return fmt.Errorf("id requis")
		}
		if updName == "" && updEmail == "" {
			return fmt.Errorf("rien à mettre à jour")
		}
		if err := Store.Update(updID, updName, updEmail); err != nil {
			return err
		}
		fmt.Println("Contact mis à jour")
		return nil
	},
}

func init() {
	updateCmd.Flags().IntVar(&updID, "id", 0, "ID du contact (requis)")
	updateCmd.Flags().StringVar(&updName, "name", "", "Nouveau nom")
	updateCmd.Flags().StringVar(&updEmail, "email", "", "Nouvel email")
	RootCmd.AddCommand(updateCmd)
}
