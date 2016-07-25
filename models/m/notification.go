package m

import ()
import "github.com/jinzhu/gorm"

type Notification struct {
	gorm.Model
	User        User
	RelatedUser User  `gorm:"ForeignKey:RelatedID"`
	UserID      uint
	RelatedID   uint
	Title       string
	Subject     string
	Content     string
	IsRead      bool
}
