package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vault-utils",
	Short: "Yet Another Vault Utility",
	Long:  `Yet Another Vault Utility`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.vault-utils.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbosity")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.SetDefault("verbose", false)

	rootCmd.PersistentFlags().StringP("addr", "s", "", "Vault address")
	viper.BindPFlag("addr", rootCmd.PersistentFlags().Lookup("addr"))
	rootCmd.PersistentFlags().StringP("token", "t", "", "Vault token")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	rootCmd.PersistentFlags().StringP("engine", "e", "", "Vault kv comma separated engine names")
	viper.BindPFlag("engine", rootCmd.PersistentFlags().Lookup("engine"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in working directory, and then home directory with name ".cobra-app" (without extension).
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".vault-utils")
	}

	viper.AutomaticEnv() // read in environment variables that match

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No config file found! Counting on flags!")
	}

	verbose = viper.GetBool("verbose")

	if verbose {
		// If a config file is found, read it in.
		fmt.Fprintln(os.Stderr, "Using config:", viper.ConfigFileUsed())

		fmt.Println("--- Configuration ---")
		for s, i := range viper.AllSettings() {
			fmt.Printf("\t%s = %s\n", s, i)
		}
		fmt.Println("---")
	}

}
