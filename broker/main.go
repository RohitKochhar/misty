package main

import (
	"fmt"
	"os"

	"rohitsingh/misty-broker/service"

	"github.com/spf13/viper"
)

// main initializes the configuration using viper
// and then calls the main logic of the broker
func main() {
	// Load the configuration from viper
	if err := initConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := service.Run(); err != nil {
		fmt.Printf("error while running misty broker: %q", err)
		os.Exit(1)
	}
}

func initConfig() error {
	// The default cfgFile should be placed at $(pwd)/.misty.yaml
	cfgFilepath, err := os.Getwd()
	if err != nil {
		return err
	}
	viper.AddConfigPath(cfgFilepath)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".misty")
	// Overwrite file configuration values with env values if found
	viper.AutomaticEnv()
	// If we find a configuration file, read it in
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error while reading variables from config file")
	}
	fmt.Printf("using configuration file: %s\n", viper.ConfigFileUsed())
	return nil
}
