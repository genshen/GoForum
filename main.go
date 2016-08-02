package main

import (
	_ "gensh.me/goforum/routers"
	"gensh.me/goforum/models/database"
	"github.com/astaxie/beego"
)

func init() {
	database.InitDB()
}

func main() {
	beego.Run()
}