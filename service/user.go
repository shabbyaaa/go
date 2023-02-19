package service

import (
	"encoding/json"
	"fmt"
	"go/models"
	"strconv"

	"github.com/asaskevich/govalidator"
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

// CreateUser
// @Summary 新增用户
// @Tags user
// @Produce  json
// @Param data body dao.CreateUser{} true "参数"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	data, _ := c.GetRawData()
	var body map[string]string
	_ = json.Unmarshal(data, &body)
	if body["name"] == "" || body["password"] == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名或密码不能为空！",
			"data":    user,
		})
		return
	}

	ret := models.FindUserByName(body["name"])
	if ret.Name != "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "该用户已存在！",
			"data":    user,
		})
		return
	}

	user.Name = body["name"]
	user.Password = body["password"]
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "操作成功",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags user
// @Produce  json
// @Param id query string true "id"
// @Success 200 {string} json{"code", "message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "操作成功",
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags user
// @Produce  json
// @Param date body dao.UpdateUser true "参数"
// @Success 200 {string} json{"code", "message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	data, _ := c.GetRawData()
	var body map[string]string
	_ = json.Unmarshal(data, &body)
	user := models.UserBasic{}
	id, _ := strconv.Atoi(body["id"])
	user.ID = uint(id)
	user.Name = body["name"]
	user.Password = body["password"]
	user.Email = body["email"]
	user.Phone = body["phone"]

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println("err", err)
		c.JSON(200, gin.H{
			"code": -1,
			// TODO: 返回的err是一个对象
			"message": err,
			"data":    "",
		})
		return
	}

	fmt.Println("update :", user)
	// TODO: 已被删除的用户修改无效
	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "修改成功",
		"data":    user,
	})
}
