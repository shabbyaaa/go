package models

import (
	"fmt"
	"go/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Salt          string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time `gorm:"type:timeStamp"`
	HeartbeatTime time.Time `gorm:"type:timeStamp"`
	LogoutTime    time.Time `gorm:"column:logout_time" json:"logout_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, Password: user.Password, Email: user.Email, Phone: user.Phone})
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("phone = ?", phone).First(&user)
}

func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("email = ?", email).First(&user)
}

func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ? and password = ? ", name, password).First(&user)
	// token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.Md5EncodeLower(str)
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}
