package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserId int64
	TargetId int64
	Type int 
	Media int
	Content string
	CreateTime uint64
	ReadTime uint64
	Pic string
	Url string
	Desc string
	Amount int
}