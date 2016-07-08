package routers

import (
	"./../controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("home/swipe", &controllers.HomeController{},"get:LoadSwipe")
	beego.Router("home/hot/:start([0-9]+)", &controllers.HomeController{},"get:Hot")
	beego.Router("home/category", &controllers.HomeController{},"get:Category")
	beego.Router("home/me", &controllers.HomeController{},"get:Me")

	beego.Router("account/signup",&controllers.UserController{},"get,post:SignUp")
	beego.Router("account/signin",&controllers.UserController{},"get,post:SignIn")
	beego.Router("account/signout",&controllers.UserController{},"get:SignOut")

	beego.Router("post/:id([0-9]+)",&controllers.PostController{},"get:View")
	beego.Router("post/create/jump",&controllers.PostController{},"get:CreateJump")
	beego.Router("post/create/m",&controllers.PostController{},"get:CreateMobile;post:POST_CreateMobile")
	beego.Router("post/create/upload_token",&controllers.PostController{},"get:UploadToken")

	beego.Router("comment/:id([0-9]+)/:start([0-9]+)",&controllers.CommentController{},"get:Comment")
	beego.Router("comment/add/:id([0-9]+)",&controllers.CommentController{},"post:CommentAdd")
	//beego.AutoRouter(&controllers.UserController{})
}