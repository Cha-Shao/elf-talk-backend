package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"elf-talk/control"
)

func JoinHandler(c *gin.Context) {
	var reqBody struct {
		UUID     string `json:"uuid"`
		Password string `json:"pwd"`
	}
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "need body",
		})
		return
	}

	token, err := control.Join(reqBody.UUID, reqBody.Password)
	if err != nil {
		if err.Error() == "P" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "wrong password",
			})
			return
		} else {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"token":   token,
	})
}
