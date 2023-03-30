package console

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// exitAfterFinish variable is used to indicate if program should exit immediately after consoles/commands/processes has finished
var exitAfterFinish bool = false

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Coursefan",
	Short: "Coursefan",
	Long:  "----------------------------------------------------\nCoursefan - Online course platform\n----------------------------------------------------",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}

	errs := []error{
		envFlag(),
		migrateFlag(),
	}
	for _, err := range errs {
		if err != nil {
			logrus.Fatal(err)
		}
	}

	if exitAfterFinish {
		os.Exit(0)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.standard-layout-golang.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
