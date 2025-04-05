package handlers

import "github.com/labstack/echo/v4"

// Ping godoc
//
//	@Summary		Ping endpoint
//	@Description	Returns a pong response to indicate the service is up and running
//	@Tags			health
//	@Produce		json
//	@Success		200	{string}	string	"pong"
//	@Router			/ping [get]
func Ping(c echo.Context) error {
	return c.JSON(200, "pong")
}
