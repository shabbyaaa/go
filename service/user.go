package service

import (
	"go/models"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags user
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"code":    0,
		"data":    data,
		"message": "",
	})
}
