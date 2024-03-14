package cmd

import (
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "blog-go",
	Short: "Blog system",
	Long: `A blog system, which is a backend server for blog.

- Restful API for:
	- Posts management
	- Tags management
	- Categories management
	- Comments management
`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := log.NewStdLogger(os.Stdout)
		app, err := initApp(logger)
		if err != nil {
			panic(err)
		}

		if err := app.Run(); err != nil {
			panic(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("$HOME/.config/blog-go")
		viper.AddConfigPath("./config/")

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	// set the default values
	setDefaultConfig()

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%s\n\n\n", err.Error())
		_ = rootCmd.Help()
		os.Exit(1)
	}
	fmt.Printf("======== Using config file: %s ========\n", viper.ConfigFileUsed())
}

func setDefaultConfig() {
	// Some default values
}
