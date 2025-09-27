package cmd

import (
	"fmt"
	"path/filepath"

	"minicrm/internal/storage"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var Store storage.Storer

var RootCmd = &cobra.Command{
	Use:   "minicrm",
	Short: "Mini-CRM CLI",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			viper.AddConfigPath(".")
			viper.SetConfigName("config")
		}
		_ = viper.ReadInConfig()
		viper.SetDefault("type", "memory")
		typ := viper.GetString("type")
		switch typ {
		case "memory":
			Store = storage.NewMemoryStore()
		case "json":
			path := viper.GetString("json.path")
			if path == "" {
				path = "contacts.json"
			}
			s, err := storage.NewJSONStore(path)
			if err != nil {
				return err
			}
			Store = s
		case "gorm":
			db := viper.GetString("gorm.path")
			if db == "" {
				db = "contacts.db"
			}
			s, err := storage.NewGORMStore(db)
			if err != nil {
				return err
			}
			Store = s
		default:
			return fmt.Errorf("unknown store type: %s", typ)
		}
		return nil
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if Store != nil {
			_ = Store.Close()
		}
	},
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	defaultCfg := filepath.Join(".", "config.yaml")
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", defaultCfg, "config file (default ./config.yaml)")
}
