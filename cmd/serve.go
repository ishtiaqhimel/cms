package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ishtiaqhimel/news-portal/cms/internal/server"
	"github.com/ishtiaqhimel/news-portal/cms/internal/utils"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve serves the cms service",
	Long:  `Serve serves the cms service`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Println("serve called")
		stopCh := utils.SetupSignalHandler()
		if err := server.Serve(stopCh); err != nil {
			logrus.Errorln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().ToP("toggle", "t", false, "Help message for toggle")
}
