package main

import (
	_ "beeapp/routers"
	"github.com/astaxie/beego"
	"beeapp/models/database"
)

func init() {
	database.InitDB()
}

func main() {
	beego.Run()
}

