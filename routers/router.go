package routers

import (
	"./../controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("account/signup",&controllers.UserController{},"get,post:SignUp")
	beego.Router("account/signin",&controllers.UserController{},"get,post:SignIn")
	beego.Router("account/signout",&controllers.UserController{},"get:SignOut")

	beego.Router("post/:id([0-9]+)",&controllers.PostController{},"get:View")
	beego.Router("post/create/jump",&controllers.PostController{},"get:CreateJump")
	beego.Router("post/create/m",&controllers.PostController{},"get:CreateMobile;post:POST_CreateMobile")
	beego.Router("post/create/upload_token",&controllers.PostController{},"get:UploadToken")

	beego.Router("comment/:id([0-9]+)/:start([0-9]+)",&controllers.CommentController{},"get:Comment")
	//beego.AutoRouter(&controllers.UserController{})
}