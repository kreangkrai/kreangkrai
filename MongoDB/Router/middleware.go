package Router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Admin")
	}
}
