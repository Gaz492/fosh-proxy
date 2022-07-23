package v1

import (
	"fosh-proxy/config"
	"fosh-proxy/database"
	"fosh-proxy/relays"
	"fosh-proxy/structs"
	"github.com/gin-gonic/gin"
	"github.com/pterm/pterm"
	"net/http"
)

func Submit(c *gin.Context) {
	var data structs.EcowittData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		pterm.Error.Println("Error binding query parameters\n", err)
		errResponse(c, err, http.StatusBadRequest)
		c.Abort()
		return
	}

	if config.Cfg.SaveToDatabase {
		err := database.SaveToDB(data)
		if err != nil {
			errResponse(c, err, http.StatusInternalServerError)
			c.Abort()
			return
		}
	}
	if config.Cfg.EnableRelay {
		err := relays.RelayHandler(data)
		if err != nil {
			errResponse(c, err, http.StatusInternalServerError)
			c.Abort()
			return
		}
	}
}

func errResponse(c *gin.Context, err error, code int) {
	c.IndentedJSON(code, gin.H{
		"ok":    false,
		"error": err.Error(),
	})
}
