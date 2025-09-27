package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lister les contacts",
	RunE: func(cmd *cobra.Command, args []string) error {
		items, err := Store.GetAll()
		if err != nil {
			return err
		}
		if len(items) == 0 {
			fmt.Println("Aucun contact trouvé")
			return nil
		}
		for _, c := range items {
			fmt.Printf("ID: %d | Name: %s | Email: %s\n", c.ID, c.Name, c.Email)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
