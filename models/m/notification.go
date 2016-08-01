package m

import (
	"github.com/jinzhu/gorm"
)

type Notification struct {
	gorm.Model
	User        User
	RelatedUser User    `gorm:"ForeignKey:RelatedID"`
	UserID      uint
	RelatedID   uint    `gorm:"default:0"`
	SubjectType int
	Data     string   `gorm:"default:''"`
	IsRead      bool     `gorm:"default:false"`
}

func (Notification) TableName() string {
	return "notification"
}
