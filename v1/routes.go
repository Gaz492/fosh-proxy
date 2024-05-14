package v1

import (
	"fosh-proxy/config"
	"fosh-proxy/database"
	"fosh-proxy/relays"
	"fosh-proxy/structs"
	"github.com/gofiber/fiber/v3"
	"github.com/pterm/pterm"
	"net/http"
)

type (
	ErrorResponse struct {
		Ok    bool   `json:"ok"`
		Error string `json:"error"`
	}
)

func Submit(c fiber.Ctx) error {
	var data structs.EcowittData
	err := c.Bind().Query(&data)
	if err != nil {
		pterm.Error.Println("Error binding query parameters\n", err)
		err = errResponse(c, err, http.StatusBadRequest)
		return err
	}

	if config.Cfg.SaveToDatabase {
		err := database.InsertData(data)
		if err != nil {
			err = errResponse(c, err, http.StatusInternalServerError)
			return err
		}
	}
	if config.Cfg.EnableRelay {
		err := relays.RelayHandler(data)
		if err != nil {
			err = errResponse(c, err, http.StatusInternalServerError)
			return err
		}
	}
	return nil
}

func errResponse(c fiber.Ctx, err error, code int) error {
	c.Status(code)
	return c.JSON(ErrorResponse{
		Ok:    false,
		Error: err.Error(),
	})
}
