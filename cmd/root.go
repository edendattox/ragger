package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/edendattox/ragger/constants"
	"github.com/edendattox/ragger/utils"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	userLicense string
	rootCmd     = &cobra.Command{
		Use:   "ragger",
		Short: "ragger is a RAG system built on top of the open ai gpt-3 engine",
		Run: func(cmd *cobra.Command, args []string) {
			// Print the greeting message
			utils.PrintColoredMessage("Hi! How can I help you today?", constants.ColorRed)

			// Create a new scanner to read user input
			scanner := bufio.NewScanner(os.Stdin)

			// Keep reading input until the user exits or sends a termination signal
			for scanner.Scan() {

				// Print the user's input
			}

			// Check for any errors that may have occurred during scanning
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
			}
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.SetDefault("name", "dummy")
	viper.SetDefault("license", "apache")

}

func initConfig() {

	// Load values from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")

	}

	viper.AutomaticEnv()

	// Enable debugging
	// viper.Debug()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	apiKey := viper.GetString("openai.api_key")

	if apiKey == "" {
		fmt.Println("API key not found in config file")
		return
	}
}
