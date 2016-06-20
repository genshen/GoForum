package routers

import (
	"./../controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("account/login",&controllers.UserController{},"get:Login;post:LoginPost")
	beego.Router("account/logout",&controllers.UserController{},"get:Logout")
	//beego.AutoRouter(&controllers.UserController{})
}