package garage

import (
	"github.com/labstack/echo"
	"meridianhomeforsale/app/garage"
)

func ChangeGarageStatus(ctx echo.Context) error {

	status := ctx.Param("status")
	firstname := ctx.Param("firstname")
	lastname := ctx.Param("lastname")
	phone := ctx.Param("phone")

	if status == "open" {
		response := garage.OpenGarage(firstname, lastname, phone)
		return ctx.JSON(200, response)
	}
	if status == "close" {
		response := garage.CloseGarage(firstname, lastname, phone)
		return ctx.JSON(200, response)
	}

	return ctx.JSON(400, "Failed")
}