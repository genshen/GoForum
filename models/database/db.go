package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"fmt"
)

var DB *gorm.DB

func InitDB() {
	var err error;
	DB, err = gorm.Open(beego.AppConfig.String("db_type"), beego.AppConfig.String("db_config"))
	if err != nil {
		fmt.Println("error to connect to database")
		return;
	}
}

//func GetDB() *gorm.DB{
//	return DB
//}
