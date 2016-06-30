package m

import (
	"github.com/jinzhu/gorm"
	"time"
	"./../database"
)

const (
	STATUS_ACTIVE int = 10
)

type User struct {
	gorm.Model
	Email      string
	Tel        string
	Name       string
	Password   string  `gorm:"column:password_hash"`
	AuthKey    string  `gorm:"column:auth_key"`
	ResetToken string  `gorm:"column:password_reset_token"`
	Status     int     `gorm:"default:10"`
	Posts      []Posts `gorm:"ForeignKey:Author"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) GetUserById(id uint) {
	database.DB.First(&u, id)
}

//User         User `gorm:"ForeignKey:Author"`
type Posts struct {
	gorm.Model
	Topic        uint
	Author       uint
	Title        string
	Content      string
	IsMobile     bool // 1 for mobile,0 for desktop
	Sticky       bool  `gorm:"default:false"`
	CommentCount int   `gorm:"default:0"`
	ViewCount    int   `gorm:"default:0"`
	Visible      bool  `gorm:"default:true"`
	LastReplayAt time.Time
}

func (p *Posts) GetPostById(id string) {
	//todo string to uint
	database.DB.First(&p, id)
}