package cmd

import (
	"github.com/ishtiaqhimel/news-portal/cms/infrastructure/db"
	"github.com/ishtiaqhimel/news-portal/cms/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var (
	migrationPath string
	uri           string
)

// migrateCmd represents the migration command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run db migration",
	Long:  `Run db migration`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Println("migration called")
		if err := config.Load(); err != nil {
			logrus.Error(err)
			os.Exit(1)
		}

		if err := db.Migrate(uri, migrationPath); err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
	},
}

func init() {
	migrateCmd.Flags().StringVar(&migrationPath, "path", "/db/migrations", "migrations file path")
	migrateCmd.Flags().StringVar(&uri, "uri", "", "postgres uri")
	rootCmd.AddCommand(migrateCmd)
}
