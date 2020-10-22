package Router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kriangkrai/Mee/MongoDB/Controller"
	"github.com/kriangkrai/Mee/MongoDB/Models"
)

func Get(c *gin.Context) {
	data := Controller.ReadDoc("Sensor3")
	// c.JSON(http.StatusOK, gin.H{
	// 	"Data": data,
	// })

	c.JSON(200, data)
}

func Insert(c *gin.Context) {
	var input Models.DataModel
	err := c.ShouldBindJSON(&input)
	if err != nil {
		fmt.Println(err)
	}
	_, err = Controller.InsertDoc(input)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(200, "Insert Success")
}

func Update(c *gin.Context) {
	var input Models.DataModel
	e := c.BindJSON(&input)
	if e != nil {
		fmt.Println(e)
	}
	_, err := Controller.UpdateDoc(input)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(200, "Update Success")
}

func Delete(c *gin.Context) {
	device := c.Params.ByName("device")
	_, err := Controller.DeleteDoc(device)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(200, "Delete Success")
}
