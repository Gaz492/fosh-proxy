package database

import (
	"fosh-proxy/config"
	"fosh-proxy/structs"
	"github.com/pterm/pterm"
)

func Initialization() {
	pterm.Info.Println("Initializing database...")
	dbType := config.Cfg.DatabaseType
	switch dbType {
	case "mysql":
		pterm.Info.Println("Setting DB type to MySQL")
	case "postgresql":
		pterm.Info.Println("Setting DB type to PostgreSQL")
	case "influx1":
		pterm.Info.Println("Setting DB type to InfluxDB v1")
	case "influx2":
		pterm.Info.Println("Setting DB type to InfluxDB v2")
	}
}

func SaveToDB(wData structs.EcowittData) error {
	return nil
}
