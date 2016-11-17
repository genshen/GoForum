package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
)

var O orm.Ormer

func init() {
	//// mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	// err := orm.RegisterDriver(beego.AppConfig.String("db_type"), orm.DRMySQL)
	// if err != nil {
	//	log.Fatal("err to connect to registe database drive")
	// }
	err := orm.RegisterDataBase("default", beego.AppConfig.String("db_type"), beego.AppConfig.String("db_config"))
	if err != nil {
		log.Fatal("err to connect to database")
	}
	orm.Debug = beego.AppConfig.DefaultBool("db_debug", false)
}
