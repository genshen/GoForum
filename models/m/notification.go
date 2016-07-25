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
	TargetID    uint    `gorm:"default:0"`
	Title       string
	Subject     string
	SubjectType int
	Content     string   `gorm:"default:''"`
	IsRead      bool     `gorm:"default:false"`
}

func (Notification) TableName() string {
	return "notification"
}
