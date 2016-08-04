package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
)

var O orm.Ormer

func init() {
	err := orm.RegisterDriver(beego.AppConfig.String("db_type"), orm.DRMySQL)
	if err != nil {
		log.Fatal(err)
	}
	orm.RegisterDataBase("default", beego.AppConfig.String("db_type"), beego.AppConfig.String("db_config"))
	if err != nil {
		log.Fatal(err)
	}
	orm.Debug = beego.AppConfig.DefaultBool("db_debug", false)
}
