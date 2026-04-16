package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "evcc",
	Short: "EV Charge Controller",
	Long: `evcc is an EV charge controller and home energy management system.
It supports a wide range of wallboxes, meters, and vehicles.`,
	RunE: run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is evcc.yaml)")
	// Default verbose logging to false to reduce noise in normal operation
	rootCmd.PersistentFlags().BoolP("log", "l", false, "enable verbose logging")

	_ = viper.BindPFlag("log", rootCmd.PersistentFlags().Lookup("log"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Search for config in common locations
		viper.SetConfigName("evcc")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath("/etc/evcc")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Fprintln(os.Stderr, "error reading config:", err)
			os.Exit(1)
		}
	}
}

// run is the main entry point for the evcc command
func run(cmd *cobra.Command, args []string) error {
	fmt.Printf("evcc version %s (%s, %s)\n", version, commit, date)
	return nil
}
