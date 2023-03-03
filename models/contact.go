package models

import (
	"go/utils"

	"gorm.io/gorm"
)

// 关系
type Contact struct {
	gorm.Model
	OwnerId  uint // 谁的关系信息
	TargetId uint // 对应的谁
	Type     int  // 对应的类型 0 1 3
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}

func SearchFriend(userId uint) []UserBasic {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id = ? and type=1", userId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, uint64((v.TargetId)))
	}
	users := make([]UserBasic, 0)
	utils.DB.Where("id in ?", objIds).Find(&users)
	return users
}

func AddFriend(userId uint, targetName string) (int, string)
