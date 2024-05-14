package main

import (
	"flag"
	"fosh-proxy/config"
	v1Router "fosh-proxy/v1"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
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
		//gin.SetMode(gin.DebugMode)
	} else {
		//gin.SetMode(gin.ReleaseMode)
	}

	pterm.Info.Println("Initializing config")
	config.Initialize(configLocation)
	pterm.Success.Println("Config initialized")

	//if config.Cfg.SaveToDatabase {
	//	database.Initialization()
	//}
}

func main() {
	//defer database.Conn.Close()
	app := fiber.New()

	app.Use(logger.New())

	v1 := app.Group("/v1")
	{
		v1.Get("/submit", v1Router.Submit)
	}

	pterm.Fatal.Println(app.Listen(":8080"))
}
