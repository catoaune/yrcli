package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	toml "github.com/BurntSushi/toml"
	netatmo "github.com/exzz/netatmo-api-go"
)

// Command line flag
var fConfig = flag.String("f", "", "Configuration file")

// NetatmoConfig contains the API credentials
type NetatmoConfig struct {
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

var config NetatmoConfig

func main() {

	// Parse command line flags
	flag.Parse()
	if *fConfig == "" {
		fmt.Printf("Missing required argument -f\n")
		os.Exit(0)
	}

	// Read API credentials from file
	if _, err := toml.DecodeFile(*fConfig, &config); err != nil {
		fmt.Printf("Cannot parse config file: %s\n", err)
		os.Exit(1)
	}

	n, err := netatmo.NewClient(netatmo.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Username:     config.Username,
		Password:     config.Password,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dc, err := n.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ct := time.Now().UTC().Unix()

	for _, station := range dc.Stations() {
		fmt.Printf("Station : %s\n", station.StationName)

		for _, module := range station.Modules() {
			fmt.Printf("\tModule : %s\n", module.ModuleName)
			ts, data := module.Data()
			for dataName, value := range data {
				if module.ModuleName == "Outdoor module" {
					if dataName == "Temperature" {
						fmt.Printf("\t\t%s : %v (updated %ds ago)\n", dataName, value, ct-ts)
					}
				}
				if module.ModuleName == "Indoor" {
					if dataName == "Temperature" {
						fmt.Printf("\t\t%s : %v (updated %ds ago)\n", dataName, value, ct-ts)
					}
				}
			}
		}
	}
}
