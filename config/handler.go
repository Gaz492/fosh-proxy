package config

import (
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
	"os"
)

var Cfg Config

func Initialize(configLocation *string) {
	cF, err := os.Open(*configLocation)
	if err != nil {
		pterm.Fatal.Println("Unable to open config file\n", err)
	}
	defer cF.Close()

	decoder := yaml.NewDecoder(cF)
	err = decoder.Decode(&Cfg)
	if err != nil {
		pterm.Fatal.Println("Unable to decode config\n", err)
	}
}
