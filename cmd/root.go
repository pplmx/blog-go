package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
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
		app := NewApp()
		// log.Fatal(app.Listen(":3000"))

		// Listen from a different goroutine to allow for graceful shutdown
		go func() {
			if err := app.Listen(":3000"); err != nil {
				log.Panic(err)
			}
		}()

		c := make(chan os.Signal, 1)                    // Create a channel to signify a signal being sent
		signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

		_ = <-c // This blocks the main thread until an interrupt is received
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()

		fmt.Println("Running cleanup tasks...")

		// Your cleanup tasks go here
		// db.Close()
		// redisConn.Close()
		fmt.Println("Fiber was successful shutdown.")
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
