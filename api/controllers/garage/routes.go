package garage

import "github.com/labstack/echo"

func Routes(e *echo.Echo) {

	g := e.Group("/home/garage")

	// Countries Group Handlers
	g.GET("/:status/:firstname/:lastname/:phone", ChangeGarageStatus)

}
