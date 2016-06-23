package routers

import (
	"./../controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("account/signup",&controllers.UserController{},"get:SignUp;post:SignUpPost")
	beego.Router("account/signin",&controllers.UserController{},"get:SignIn;post:SignInPost")
	beego.Router("account/signout",&controllers.UserController{},"get:SignOut")
	//beego.AutoRouter(&controllers.UserController{})
}