package main

import (
	_ "gensh.me/goforum/routers"
	_"gensh.me/goforum/models/database" //init database
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}