package main

import (
	_ "./routers"
	"./models/database"
	"github.com/astaxie/beego"
)

func init() {
	database.InitDB()
}

func main() {
	beego.Run()
}