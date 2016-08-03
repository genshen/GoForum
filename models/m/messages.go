package m

import (
	"github.com/jinzhu/gorm"
)

type BaseMessage struct {
	gorm.Model
	User        User
	RelatedUser User    `gorm:"ForeignKey:RelatedID"`
	UserID      uint
	RelatedID   uint    `gorm:"default:0"`
	SubjectType int
	IsRead      bool     `gorm:"default:false"`
}

type Notification struct {
	BaseMessage
	Data string   `gorm:"default:''"`
}

func (Notification) TableName() string {
	return "notification"
}

type PostMessage struct {
	BaseMessage
	RelatedUsername string   `gorm:"default:''"`
	PostID          uint
	PostTitle       string
	Summary         string `gorm:"size:255"`
	Quote           string `gorm:"size:255"`
}

func (PostMessage) TableName() string {
	return "post_message"
}
