package Router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kriangkrai/Mee/MongoDB/Controller"
)

func SetupRouter() (*gin.Engine, string) {
	Controller.Connect()

	r := gin.Default()
	r.Use(Middleware())
	r.GET("/get", Get)
	r.POST("/insert", Insert)
	r.PUT("/update", Update)
	r.DELETE("/delete/:device", Delete)
	port := "8080"
	if os.Getenv("ASPNETCORE_PORT") != "" { // get enviroment variable that set by ACNM
		port = os.Getenv("ASPNETCORE_PORT")
	}

	return r, port
}
