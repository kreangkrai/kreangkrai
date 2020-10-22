package main

import (
	"github.com/kriangkrai/Mee/MongoDB/Controller"
	"github.com/kriangkrai/Mee/MongoDB/Router"
)

func main() {

	r, port := Router.SetupRouter()
	r.Run(":" + port)
	Controller.Disconnect()
}
