package main

import (
	"flag"
	"fosh-proxy/config"
	"fosh-proxy/database"
	v1Router "fosh-proxy/v1"
	"github.com/gin-gonic/gin"
	"github.com/pterm/pterm"
)

var (
	configLocation *string
	debug          *bool
)

func init() {
	configLocation = flag.String("c", "config.yml", "Config file location")
	debug = flag.Bool("debug", false, "Enable debug logging")

	flag.Parse()

	if *debug {
		pterm.EnableDebugMessages()
		pterm.Debug.Println("Debug logging enabled")
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	pterm.Info.Println("Initializing config")
	config.Initialize(configLocation)
	pterm.Success.Println("Config initialized")

	if config.Cfg.SaveToDatabase {
		database.Initialization()
	}
}

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/submit", v1Router.Submit)
	}

	err := r.Run()
	if err != nil {
		pterm.Error.Println("Error starting server", err)
	}
}
